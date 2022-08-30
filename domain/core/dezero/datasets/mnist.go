package datasets

import (
	"image/color"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/util"
	"github.com/fcfcqloow/go-advance/log"
	"github.com/petar/GoMNIST"
)

const (
	TRAIN_URL      = "http://yann.lecun.com/exdb/mnist/train-images-idx3-ubyte.gz"
	LABEL_URL      = "http://yann.lecun.com/exdb/mnist/train-labels-idx1-ubyte.gz"
	TEST_TRAIN_URL = "http://yann.lecun.com/exdb/mnist/t10k-images-idx3-ubyte.gz"
	TEST_LABEL_URL = "http://yann.lecun.com/exdb/mnist/t10k-labels-idx1-ubyte.gz"

	LOCAL_DATA_PATH        = "data"
	LOCAL_TRAIN_PATH       = LOCAL_DATA_PATH + "/train-images-idx3-ubyte.gz"
	LOCAL_LABELS_PATH      = LOCAL_DATA_PATH + "/train-labels-idx1-ubyte.gz"
	LOCAL_TEST_TRAIN_PATH  = LOCAL_DATA_PATH + "/t10k-images-idx3-ubyte.gz"
	LOCAL_TEST_LABELS_PATH = LOCAL_DATA_PATH + "/t10k-labels-idx1-ubyte.gz"
)

type mnist struct{ dz.Dataset }

func NewMnist(options ...dz.DatasetOption) dz.Dataset {
	log.Info("Loading mnist data")
	defer log.Info("Loaded mnist data")

	downloadMnist()
	opt := dz.ApplyDataSetOpt(options...)

	trainSet, testSet, err := GoMNIST.Load(LOCAL_DATA_PATH)
	if err != nil {
		panic(err)
	}

	if opt.Train != nil && *opt.Train {
		trainData, trainLabels := loadMatrix(trainSet)
		return &mnist{Dataset: dz.ExtendsDataset(trainData, trainLabels, func() {}, options...)}
	}

	testData, testLabels := loadMatrix(testSet)
	return &mnist{Dataset: dz.ExtendsDataset(testData, testLabels, func() {}, options...)}

}

func downloadMnist() error {
	urls := []string{TRAIN_URL, LABEL_URL, TEST_TRAIN_URL, TEST_LABEL_URL}
	localPaths := []string{LOCAL_TRAIN_PATH, LOCAL_LABELS_PATH, LOCAL_TEST_TRAIN_PATH, LOCAL_TEST_LABELS_PATH}
	if !util.FileExists(LOCAL_TRAIN_PATH) ||
		!util.FileExists(LOCAL_LABELS_PATH) ||
		!util.FileExists(LOCAL_TEST_TRAIN_PATH) ||
		!util.FileExists(LOCAL_TEST_LABELS_PATH) {
		if !util.FileExists(LOCAL_DATA_PATH) {
			if err := os.Mkdir(LOCAL_DATA_PATH, os.ModePerm); err != nil {
				return err
			}
		}

		for i := 0; i < len(urls); i++ {
			log.Info("Download", urls[i], "...")

			response, err := http.Get(urls[i])
			if err != nil {
				return err
			}

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			log.Info("Save", localPaths[i])
			if err := ioutil.WriteFile(localPaths[i], body, os.ModePerm); err != nil {
				return err
			}

		}
	}
	return nil
}

func loadMatrix(s *GoMNIST.Set) (data, labels core.Matrix) {
	mat := core.NewMat(core.Shape{R: len(s.Images), C: s.NRow * s.NRow})
	lMat := core.NewMat(core.Shape{R: len(s.Labels), C: 1})

	for i := 0; i < len(s.Images); i++ {
		i := i
		img, label := s.Get(i)
		b := img.Bounds()
		// p := plot.New()
		// p.Add(plotter.NewImage(img, 0, 0, 200, 200))
		// err := p.Save(5*vg.Centimeter, 5*vg.Centimeter, "mnist/"+cnv.MustStr(label)+"_"+cnv.MustStr(i)+".png")
		// if err != nil {
		// 	log.Debug("error saving image plot: %v\n", err)
		// }
		row := make([]float64, 0, s.NRow*s.NRow)
		for x := 0; x < b.Max.Y; x++ {
			for y := 0; y < b.Max.X; y++ {
				row = append(row, float64(img.At(x, y).(color.Gray).Y))
			}
		}

		mat.SetRow(i, row)
		lMat.SetRow(i, []float64{float64(label)})
	}

	return mat, lMat
}
