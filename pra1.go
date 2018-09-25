package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var dia int
var mes int
var ano int
var dataVenda = make([]string, 0)
var codCliente = make([]int, 0)
var codFunc = make([]int,0)
var reais  int
var cents  int
var valor = make([]string,0)
var itens = [5]string{"Copo", "Garrafa", "Carteira", "Sapato", "Secador"}
var itensComprados = make([]string,0)
var qtdDisc = make([]int, 0)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func escreveValores(id int, f *os.File) {
	w := bufio.NewWriter(f)
	w.WriteString(strconv.Itoa(codCliente[id]))
	w.WriteString(";")
	w.WriteString(strconv.Itoa(codFunc[id]))
	w.WriteString(";")
	w.WriteString(dataVenda[id])
    w.WriteString(";")
    w.WriteString(valor[id])
    w.WriteString(";")
    w.WriteString(itensComprados[id])
	w.WriteString("\n")
	w.Flush()
}

func ordena(f *os.File) {
	n, err := f.Stat()
	if err != nil {
		panic(err)
	}
	tamanho := n.Size()

	// separa ao meio, se o tamanho é maior que a memoria separa denovo
	// após isso roda a equação e ve quantas vezes sera feia a ordenação
	// http://www.decom.ufop.br/guilherme/BCC203/geral/ed2_ordenacao-externa.pdf
	// não sei se isso realmente funciona mas ta ai
}

func criaValores(id int) {
	//data de venda
	dia = rand.Intn(28) + 1
	mes = rand.Intn(12) + 1
	ano = rand.Intn(8) + 2010
	dataVenda = append(dataVenda, strconv.Itoa(dia)+"/"+strconv.Itoa(mes)+"/"+strconv.Itoa(ano))

	//codigo do funcionario
    codFunc = append(codFunc, id + 1)

    //codigo do Cliente
    codCliente = append(codCliente,id + 1)



    //Valor de Venda
    reais = rand.Intn(99) + 1
    cents = rand.Intn(99) + 1
    valor = append(valor,"R$ "+strconv.Itoa(reais)+","+strconv.Itoa(cents))

	//Itens
	idItem1 := rand.Intn(5)
	idItem2 := rand.Intn(5)
	itensComprados = append(itensComprados, itens[idItem1]+" "+itens[idItem2])

}

func leArquivo(f *os.File, qtd int) {
	scanner := bufio.NewScanner(f)
	reader := bufio.NewReader(os.Stdin)
	cont := 0
	for scanner.Scan() {
		if cont%qtd == 0 {
			reader.ReadBytes('\n')
		}
		if cont == qtd {
			reader.ReadBytes('\n')
		}
		fmt.Println(scanner.Text())
		cont++
	}
}

func main() {

	t := time.Now()

	var escolha int
	fmt.Println("Manipulação de arquivos \n1 - Criar \n2 - Ler")
    fmt.Scan(&escolha)
    // testa se foi escolhido uma opção valida
	if escolha != 1 && escolha != 2 {
		fmt.Println("ERRO! \nopção invalida!")
	} else if escolha == 1 {
        // cria o arquivo
		f, err := os.Create("arquivo.txt")
		check(err)

		defer f.Close()

		// pergunta se será gerado por quantidade de registros ou tamanho do arquivo
		var tipo int
		fmt.Println("Como você deseja criar o arquivo?:\n1 - Quantidade\n2 - Tamanho")
		fmt.Scan(&tipo)

		// testa se foi escolhido uma opção valida
		if tipo != 1 && tipo != 2 {
			fmt.Println("ERRO!\nopção invalida!")
		} else if tipo == 1 {
            // se escolhido por quantidade, questiona quantos registros serão criados
			fmt.Println("Quantos registros você deseja criar?")
			var qtdReg int
			fmt.Scan(&qtdReg)
            // Questiona quantas vezes será gerado esses registros
			fmt.Println("Quantas vezes deseja criar esses registros?")
			var qtdVez int
			fmt.Scan(&qtdVez)

			t = time.Now()

			for j := 0; j < qtdVez; j++ {
				for i := 0; i < qtdReg; i++ {
                    // cria os valores e escreve no arquivo
					criaValores(i)
					escreveValores(i, f)
                }
                // mensagem de acompanhamento no terminal
				fmt.Println("Criando ", qtdReg, " registros")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadBytes('\n')
			}
		} else if tipo == 2 {
			tam, err := f.Stat()
			check(err)
            // se escolhido por tamanho, pede para informar o tamanho em bytes
			fmt.Println("Qual o tamanho do arquivo em bytes que você deseja criar?")
			var sizeReg int64
			fmt.Scan(&sizeReg)
			cont := 0


            t = time.Now()

            // cria os valores e escreve
            // ele da um for considerando o tamanho do arquivo tam.size e o tamanho informado sizeReg
			for tam.Size() < sizeReg {
				tam, _ = f.Stat()
				criaValores(cont)
				escreveValores(cont, f)
				cont++
			}
		}
		f.Close()
	} else if escolha == 2 {
        // leitura

		f, err := os.Open("arquivo.txt")
		check(err)

		defer f.Close()
		ordena(f)

        var qtdLer int
        // informa quantos registros serão lidos
		fmt.Println("Deseja ler quantos registros por vez?")
		fmt.Scan(&qtdLer)

        t = time.Now()

		leArquivo(f, qtdLer)
		f.Close()
	}
    d := time.Now()
    // tempo de execução
	elapsed := d.Sub(t)
	fmt.Printf("%s", elapsed)

}
