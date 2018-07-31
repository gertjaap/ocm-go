package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

var GPUType string
var statusChan chan string
var panicChan chan error
var switchPageChan chan string
var stopMiningChan chan bool
var hashrateChan chan int

type ViewModel struct {
	core.QObject

	_ string `property:"initStatus"`
	_ string `property:"activePage"`
}

func main() {
	fmt.Println("Vertcoin OCM Starting up...")
	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using QML
	gui.NewQGuiApplication(len(os.Args), os.Args)

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Default")

	// create the qml application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	viewModel := NewViewModel(nil)
	viewModel.SetInitStatus("Initializing...")
	viewModel.SetActivePage("init")
	engine.RootContext().SetContextProperty("viewModel", viewModel)

	initListeners(viewModel)

	statusChan <- "Initializing..."

	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	go setup()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	gui.QGuiApplication_Exec()
}

func initListeners(viewModel *ViewModel) {
	statusChan = make(chan string)
	switchPageChan = make(chan string)
	/*
		panicChan = make(chan error)
		stopMiningChan = make(chan bool)
		hashrateChan = make(chan int)*/

	go initStatusListener(viewModel)
	go switchPageListener(viewModel)
	/*go panicListener(view)

	go hashrateListener(view)
	go httpListener(view)*/
}

func initStatusListener(viewModel *ViewModel) {
	for {
		status := <-statusChan
		viewModel.SetInitStatus(status)
	}
}

/*func panicListener(view *webengine.QWebEngineView) {
	for {
		err := <-panicChan
		errString := strings.Replace(err.Error(), "'", "''", -1)
		script := fmt.Sprintf("setPanic('%s')", errString)
		view.Page().RunJavaScript(script)
	}
}

func hashrateListener(view *webengine.QWebEngineView) {
	for {
		hashRate := <-hashrateChan
		script := fmt.Sprintf("setHashrate('%d')", hashRate)
		view.Page().RunJavaScript(script)
	}
}
*/

func switchPageListener(viewModel *ViewModel) {
	for {
		nextPage := <-switchPageChan
		viewModel.SetActivePage(nextPage)
	}
}

func setup() {
	statusChan <- "Checking for miners dir"
	if _, err := os.Stat("./miners"); os.IsNotExist(err) {
		err := os.Mkdir("./miners", 0755)
		if err != nil {
			panic(err)
		}
	}

	gpuVendor := GetGPU()
	switch {
	case strings.Contains(gpuVendor, "Radeon"):
		if _, err := os.Stat("./miners/AMD"); os.IsNotExist(err) {
			statusChan <- "Downloading AMD miner"
			err := DownloadFile("https://github.com/CryptoGraphics/lyclMiner/releases/download/untagged-95777e4326ae4e5ccdb5/lyclMiner015.zip", "./miners/AMD.zip")
			if err != nil {
				panicChan <- err
				return
			}

			err = UnzipFile("./miners/AMD.zip", "./miners/AMD")
			if err != nil {
				panicChan <- err
				return
			}

			err = os.Remove("./miners/AMD.zip")
			if err != nil {
				panicChan <- err
				return
			}
		}

		err := os.Chdir("./miners/AMD/lyclMiner015")
		if err != nil {
			panicChan <- err
			return
		}

		if _, err := os.Stat("lycl.conf"); err == nil {
			err := os.Remove("lycl.conf")
			if err != nil {
				panicChan <- err
				return
			}
		}

		cmd := exec.Command("lyclMiner.exe", "-g", "lycl.conf")
		err = cmd.Run()
		if err != nil {
			panicChan <- err
			return
		}

		err = ReplaceInFile("lycl.conf", "stratum+tcp://example.com:port", "stratum+tcp://p2proxy.vertcoin.org:9171")
		if err != nil {
			panicChan <- err
			return
		}

		GPUType = "Radeon"
	case strings.Contains(gpuVendor, "NVIDIA"):
		if _, err := os.Stat("./miners/NVIDIA"); os.IsNotExist(err) {
			statusChan <- "Downloading NVIDIA miner"
			err := DownloadFile("https://vtconline.org/downloads/ccminer.zip", "./miners/NVIDIA.zip")
			if err != nil {
				panicChan <- err
				return
			}

			statusChan <- "Unpacking NVIDIA miner"
			err = UnzipFile("./miners/NVIDIA.zip", "./miners/NVIDIA")
			if err != nil {
				panicChan <- err
				return
			}

			err = os.Remove("./miners/NVIDIA.zip")
			if err != nil {
				panicChan <- err
				return
			}
		}

		GPUType = "NVIDIA"
	default:
		panicChan <- fmt.Errorf("Neither AMD or nVidia GPU found")
		return
	}

	switchPageChan <- "start"
}

func startMining(address string) {
	var cmd *exec.Cmd
	var err error
	var stdout io.ReadCloser
	if GPUType == "NVIDIA" {
		cmd, stdout, err = startNVIDIA(address)
	} else {
		cmd, stdout, err = startAMD(address)
	}
	if err != nil {
		panicChan <- err
		return
	}
	go func() {
		r := bufio.NewReader(stdout)
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					panicChan <- err
				}
			}

			if strings.Contains(line, "MH/s") {
				hashRate := 100000
				hashrateChan <- hashRate
			}

		}
	}()

	go func() {
		stop := <-stopMiningChan
		if stop {
			if cmd != nil {
				cmd.Process.Kill()
			}
		}
	}()
}

func stopMining() {
	stopMiningChan <- true
}

func startAMD(address string) (*exec.Cmd, io.ReadCloser, error) {
	err := ReplaceInFile("lycl.conf", "user", address)
	if err != nil {
		return nil, nil, err
	}

	cmd := exec.Command("lyclMiner.exe", "lycl.conf")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}

	return cmd, stdout, nil
}

func startNVIDIA(address string) (*exec.Cmd, io.ReadCloser, error) {
	cmd := exec.Command("./miners/NVIDIA/ccminer-x64.exe", "-a", "lyra2v2", "-o", "stratum+tcp://p2proxy.vertcoin.org:9171", "-u", address, "-p", "x")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}

	return cmd, stdout, nil
}

/*
func mainWindow(GPUType string) {
	ui.QueueMain(func() {
		var cmd *exec.Cmd
		input := ui.NewEntry()
		button := ui.NewButton("Mine")
		status := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("VTC address:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(status, false)
		mWindow := ui.NewWindow("OCM-go", 100, 50, false)
		mWindow.SetMargined(true)
		mWindow.SetChild(box)

		button.OnClicked(func(*ui.Button) {
			var err error
			var stdout io.ReadCloser
			if GPUType == "NVIDIA" {
				cmd, stdout, err = startNVIDIA(input.Text())
			} else {
				cmd, stdout, err = startAMD(input.Text())
			}

			if err != nil {
				panic(err)
			}

			status.SetText("Started mining")

			button.Disable()

			minerOutputWindow := ui.NewWindow("OCM-go", 100, 200, false)
			minerBox := ui.NewVerticalBox()
			minerOutputWindow.SetMargined(true)
			minerOutputWindow.SetChild(minerBox)
			minerOutputWindow.Show()

			go func() {
				r := bufio.NewReader(stdout)
				n := 0
				for {
					line, err := r.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						} else {
							panic(err)
						}
					}

					ui.QueueMain(func() {
						if n >= 20 {
							minerBox.Delete(0)
							n--
						}

						n++
						minerBox.Append(ui.NewLabel(line), false)
					})
				}

				ui.QueueMain(func() {
					button.Enable()
				})
			}()
		})

		mWindow.OnClosing(func(*ui.Window) bool {
			if cmd != nil {
				cmd.Process.Kill()
			}

			ui.Quit()
			return true
		})
		mWindow.Show()
	})
}
*/
