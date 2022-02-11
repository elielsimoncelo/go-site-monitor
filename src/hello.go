package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const totalIteracoes = 5
const intervaloEspera = 5
const tempoEspera = time.Second
const arquivoParaLogDoMonitoramento = "monitoramento.log"
const arquivoComSitesParaMonitoramento = "sites.txt"

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			saiDoPrograma()
		default:
			fmt.Println("Nao conheco este comando...")
			saiDoProgramaComErro()
		}
	}
}

func exibeIntroducao() {
	var nome string = "Jose"
	var idade = 18 // inferencia de tipo
	//var versao float32 = 1.1
	//var versao = 1.1 // inferencia de tipo / variavel
	versao := 1.1 // auto inferencia de tipo e declaracao de variavel

	fmt.Println("Ola", nome, "sua idade e", idade)
	fmt.Println("Este programa esta na versao: ", versao)
	fmt.Println("O tipo da variavel versao e", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("[1] Iniciar o monitoramento")
	fmt.Println("[2] Exibir logs")
	fmt.Println("[0] sair do programa")
}

func leComando() int {
	var comandoLido int

	//fmt.Scanf("%d", &comando) // definindo com o modificador %d para definir o tipo de entrada
	fmt.Scan(&comandoLido) // sem a necessidade do modificador auto inferencia do tipo
	//fmt.Println("O endereco da variavel comando e", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitoramento iniciado...")
	fmt.Println("-------------------------")

	sites := leSitesDoArquivo()

	fmt.Println("Sites: ", sites)

	for i := 0; i < totalIteracoes; i++ {
		for index, site := range sites {
			fmt.Println("Testando site", index, ":", site)
			testaSite(site)
		}

		time.Sleep(intervaloEspera * tempoEspera)
		fmt.Println("")
	}

	fmt.Println("")
}

func leSitesDoArquivo() []string {
	var sites []string

	_, executavel, _, _ := runtime.Caller(0)
	caminho := path.Join(path.Dir(executavel), arquivoComSitesParaMonitoramento)

	//arquivo, err := ioutil.ReadFile(caminho)
	//fmt.Println(string(arquivo))

	arquivo, err := os.Open(caminho)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if err == io.EOF {
			break
		}

		if linha == "" {
			continue
		}

		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resp.StatusCode)
		registraLog(site, false)
	}
}

func registraLog(site string, status bool) {
	_, executavel, _, _ := runtime.Caller(0)
	caminho := path.Join(path.Dir(executavel), arquivoParaLogDoMonitoramento)

	arquivo, err := os.OpenFile(caminho, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " | online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	_, executavel, _, _ := runtime.Caller(0)
	caminho := path.Join(path.Dir(executavel), arquivoParaLogDoMonitoramento)

	arquivo, err := ioutil.ReadFile(caminho)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

func saiDoPrograma() {
	os.Exit(0)
}

func saiDoProgramaComErro() {
	os.Exit(-1)
}
