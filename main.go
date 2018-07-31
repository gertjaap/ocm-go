package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/widgets"
)

var GPUType string
var statusChan chan string
var panicChan chan error
var switchPageChan chan string
var stopMiningChan chan bool
var hashrateChan chan int

func main() {
	fmt.Println("Vertcoin OCM Starting up...")
	widgets.NewQApplication(len(os.Args), os.Args)
	webengine.QtWebEngine_Initialize()
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Vertcoin One Click Miner")

	var view = webengine.NewQWebEngineView(nil)
	view.Load(core.NewQUrl3("qrc:/web/index.html", 0))

	window.SetCentralWidget(view)
	window.Show()

	view.ConnectLoadFinished(func(vbo bool) {
		if vbo {
			initListeners(view)
			go setup()
		}
	})

	widgets.QApplication_Exec()
}

func initListeners(view *webengine.QWebEngineView) {
	statusChan = make(chan string)
	switchPageChan = make(chan string)
	panicChan = make(chan error)
	stopMiningChan = make(chan bool)
	hashrateChan = make(chan int)

	go initStatusListener(view)
	go panicListener(view)
	go switchPageListener(view)
	go hashrateListener(view)
	go httpListener(view)
}

func initStatusListener(view *webengine.QWebEngineView) {
	for {
		status := <-statusChan
		status = strings.Replace(status, "'", "''", -1)
		script := fmt.Sprintf("setStatus('%s')", status)
		view.Page().RunJavaScript(script)
	}
}

func panicListener(view *webengine.QWebEngineView) {
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

func switchPageListener(view *webengine.QWebEngineView) {
	for {
		nextPage := <-switchPageChan
		script := fmt.Sprintf("showPage('%s')", nextPage)
		view.Page().RunJavaScript(script)
	}
}

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL)
		headers := w.Header()
		headers.Add("Access-Control-Allow-Origin", "qrc://")
		headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")
		headers.Add("Access-Control-Allow-Methods", "GET,OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			h(w, r)
		}
	}
}

func httpListener(view *webengine.QWebEngineView) {
	http.HandleFunc("/startmining", corsHandler(func(resp http.ResponseWriter, req *http.Request) {
		startMining(req.URL.Query().Get("address"))
	}))
	http.HandleFunc("/stopmining", corsHandler(func(resp http.ResponseWriter, req *http.Request) {
		stopMining()
	}))
	port := 32001
	listenString := fmt.Sprintf("127.0.0.1:%d", port)
	view.Page().RunJavaScript(fmt.Sprintf("setHost('http://%s')", listenString))
	fmt.Printf("Listening on %s\n", listenString)
	err := http.ListenAndServe(listenString, nil)

	for err != nil {
		fmt.Println(err.Error())
		port++
		listenString := fmt.Sprintf("127.0.0.1:%d", port)
		fmt.Printf("Listening on %s\n", listenString)
		view.Page().RunJavaScript(fmt.Sprintf("setHost('http://%s')", listenString))
		err = http.ListenAndServe(listenString, nil)
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

	switchPageChan <- "mainPage"
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
