<template>
    <div>
        <section class="todoapp" v-cloak>
            <header class="header">
                <h1>Teoria de la Informacion</h1>
                <div class="buttons">
                    <ul class="filters">
                        <li>
                            <label>
                                <input class="file-name" autofocus autocomplete="off" placeholder="Path al archivo">
                            </label>
                        </li>
                        <li>
                            <a @click="loadFile">Abrir</a>
                        </li>
                    </ul>
                </div>
            </header>
            <section class="main">
                <table>
                    <tr >
                        <td class="opciones">
                            <ul class="hamming">
                                <li>
                                    <a @click="doHamming">
                                        Codificar con Hamming
                                    </a>
                                </li>
                                <li>
                                    <a @click="doDeHammingC">
                                        Decodificar Hamming
                                    </a>
                                </li>
                                <li class="aclaracion">
                                    Corrigiendo errores
                                </li>
                                <li>
                                    <a @click="doDeHammingE">
                                        Decodificar Hamming
                                    </a>
                                </li>
                                <li class="aclaracion">
                                    Sin corregir errores
                                </li>
                                <li>
                                    <a @click="doErrors">
                                        Introducir errores
                                    </a>
                                </li>
                            </ul>
                        </td>
                        <td>
                            <div class="right" :class="{used: hamming}">
                                <HammingOptions/>
                            </div>
                            <div class="right" :class="{used: deHammingC}">
                                <DeHammingCorrector/>
                            </div>
                            <div class="right" :class="{used: deHammingE}">
                                <DeHammingErrors/>
                            </div>
                            <div class="right" :class="{used: errors}">
                                <AddErrors/>
                            </div>
                            <div class="right" :class="{used: huffman}">
                                <Huffman/>
                            </div>
                            <div class="right" :class="{used: deHuffman}">
                                <DeHuffman/>
                            </div>
                        </td>
                    </tr>
                    <tr>
                        <td class="opciones">
                            <ul class="huffman">
                                <li>
                                    <a @click="doHuffman">
                                        Comprimir con Huffman
                                    </a>
                                </li>
                                <li>
                                    <a @click="doDeHuffman">
                                        Descomprimir
                                    </a>
                                </li>
                            </ul>
                        </td>
                    </tr>
                </table>
            </section>
        </section>
    </div>
</template>

<script>
    import HammingOptions from "./components/HammingOptions";
    import Huffman from "./components/Huffman";
    import DeHammingCorrector from "./components/DeHammingCorrector";
    import DeHammingErrors from "./components/DeHammingUncorrected";
    import AddErrors from "./components/AddErrors";
    import DeHuffman from "./components/DeHuffman";
    export default {
        name: "app",
        components: {DeHuffman, AddErrors, DeHammingErrors, DeHammingCorrector, HammingOptions, Huffman},
        data() {
            return{
                hamming: false,
                deHammingC: false,
                deHammingE: false,
                errors: false,
                huffman: false,
                deHuffman: false,
            };
        },
        methods: {
            loadFile: function() {
                // eslint-disable-next-line no-console
                console.log("An file will be opened")
            },
            doHamming: function () {
                 this.hamming = true;
                 this.deHammingC = false;
                 this.deHammingE = false;
                 this.errors = false;
                 this.huffman = false;
                 this.deHuffman = false;
            },
            doDeHammingC: function () {
                this.hamming = false;
                this.deHammingC = true;
                this.deHammingE = false;
                this.errors = false;
                this.huffman = false;
                this.deHuffman = false;
            },
            doDeHammingE: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = true;
                this.errors = false;
                this.huffman = false;
                this.deHuffman = false;
            },
            doErrors: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.errors = true;
                this.huffman = false;
                this.deHuffman = false;
            },
            doHuffman: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.errors = false;
                this.huffman = true;
                this.deHuffman = false;
            },
            doDeHuffman: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.errors = false;
                this.huffman = false;
                this.deHuffman = true;
            },
        },
    };
</script>
<style>
    * {
        box-sizing: border-box;
        font-family: "Roboto", Helvetica, Arial, sans-serif;
    }
    /* html */
    html{
        text-align: center;
        margin-top: 20px;
    }
    /* table */
    table {
        width: 100%;
    }
    /* li */
    .filters li {
        display: inline;
    }
    /* open file */
    .filters a {
        background-color: white;
        color: black;
        border: 2px solid #4CAF50; /* Green */
        margin-left: 5px;
        padding: 8px;
        transition-duration: 400ms;
        border-radius: 8px;
    }
    .filters a:hover {
        background-color: #4CAF50; /* Green */
        color: white;
    }
    .filters a:active {
        background-color: #087B0C;
        color: black;
        border: 2px solid #4CAF50;
        border-radius: 50%;
        transition-duration: 400ms;
    }

    .opciones {
        width: 40%;
        min-width: max-content;
    }
    /* select action */
    /* Hamming */
    .hamming a {
        display: flex;
        justify-content: center;
        align-items: center;
        border: 2px solid #000000;
        border-radius: 8px;
        padding: 4px;
        width: 100%;
        height: 100%;
        background-color: #6CCA51;
    }

    ul.hamming{
        list-style-type: none;
        width: 100%;
    }

    .hamming li{
        text-align: center;
        padding: 8px;
        margin: 5px;
    }

    li.aclaracion {
        margin-top: -16px;
        font-size: 0.9em;
        color: #FA6095;
    }
    /* Huffman */
    .huffman a {
        display: flex;
        justify-content: center;
        align-items: center;
        border: 2px solid #000000;
        border-radius: 8px;
        padding: 4px;
        width: 100%;
        height: 100%;
        background-color: #A3EB6A;
    }

    ul.huffman{
        list-style-type: none;
        width: 100%;
    }

    .huffman li{
        text-align: center;
        padding: 8px;
        margin: 5px;
    }

    .huffman a:hover, .hamming a:hover {
        border: 2px solid #00ff00;
        background-color: #00bb00;
    }

    .huffman a:active, .hamming a:active {
        border: 2px solid #000000;
        background-color: #00ff00;
    }

    .right {
        width: 60%;
        display: none;
    }

    .used {
        display: inherit;
    }

</style>