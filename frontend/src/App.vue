<template >
    <div :class="{nightMode: nigth}">
        <section class="todoapp" v-cloak>
            <header class="header">
                <label>
                    <input class="toggle" type="checkbox" v-model="nigth">
                </label>
                <h1>Teoria de la Informacion</h1>
                <div class="buttons">
                    <ul class="filters" >
                        <li>
                            <label>
                                <input class="file-name" autofocus autocomplete="off" v-model="file">
                            </label>
                        </li>
                        <li>
                            <slot><input type="file" /></slot>
                            <button class="custom-upload" @click="showUpload">Abrir archivo</button>
                        </li>
                    </ul>
                </div>
            </header>
            <section class="main">
                <div
                        :class="{nightMode: nigth}"
                        style="height: 700px"
                        v-if="!file"
                >
                </div>
                <table v-if="file" :class="{nightMode: nigth}">
                    <tr >
                        <td class="opciones">
                            <ul class="hamming">
                                <li>
                                    <a
                                            @click="doHamming"
                                            v-if="
                                              extension !== 'ha1' &&
                                              extension !== 'ha2' &&
                                              extension !== 'ha3' &&
                                              extension !== 'ha4'"
                                    >
                                        Codificar con Hamming
                                    </a>
                                </li>
                                <li>
                                    <a
                                            @click="doDeHammingC"
                                            v-if="
                                              extension === 'ha1' ||
                                              extension === 'ha2' ||
                                              extension === 'ha3' ||
                                              extension === 'ha4'"
                                    >
                                        Decodificar Hamming
                                    </a>
                                </li>
                                <li
                                        class="aclaracion"
                                        v-if="
                                              extension === 'ha1' ||
                                              extension === 'ha2' ||
                                              extension === 'ha3' ||
                                              extension === 'ha4'"
                                >
                                    Corrigiendo errores
                                </li>
                                <li>
                                    <a
                                            @click="doDeHammingE"
                                            v-if="
                                              extension === 'ha1' ||
                                              extension === 'ha2' ||
                                              extension === 'ha3' ||
                                              extension === 'ha4'"
                                    >
                                        Decodificar Hamming
                                    </a>
                                </li>
                                <li
                                        class="aclaracion"
                                        v-if="
                                              extension === 'ha1' ||
                                              extension === 'ha2' ||
                                              extension === 'ha3' ||
                                              extension === 'ha4'"
                                >
                                    Sin corregir errores
                                </li>
                                <li>
                                    <a
                                            @click="doErrors"
                                            v-if="
                                              extension === 'ha1' ||
                                              extension === 'ha2' ||
                                              extension === 'ha3' ||
                                              extension === 'ha4'"
                                    >
                                        Introducir errores
                                    </a>
                                </li>
                            </ul>
                        </td>
                        <td>
                            <div class="right" :class="{used: hamming}">
                                <table class="hamming">
                                    <tr>
                                        <td>
                                            <ul>
                                                <li>
                                                    <a
                                                            @click="doHamming7"
                                                    >
                                                        Hamming 7
                                                    </a>
                                                </li>
                                                <li>
                                                    <a> Hamming 32</a>
                                                </li>
                                                <li>
                                                    <a> Hamming 1024</a>
                                                </li>
                                                <li>
                                                    <a> Hamming 32768</a>
                                                </li>
                                            </ul>
                                        </td>
                                    </tr>
                                </table>
                            </div>
                            <div class="right" :class="{used: deHammingC}">
                                <DeHammingCorrected/>
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
                            {{ this.message }}
                        </td>
                    </tr>
                    <tr>
                        <td class="opciones">
                            <ul class="huffman">
                                <li>
                                    <a @click="doHuffman" v-if="extension !== 'huf'">
                                        Comprimir con Huffman
                                    </a>
                                </li>
                                <li>
                                    <a @click="doDeHuffman" v-if="extension === 'huf'">
                                        Descomprimir
                                    </a>
                                </li>
                            </ul>
                        </td>
                    </tr>
                    <tr>
                        <td class="opciones">
                            <ul class="hamminghuffman">
                                <li>
                                    <a
                                            @click="doHH"
                                            v-if="
                                              extension !== 'hh1' &&
                                              extension !== 'hh2' &&
                                              extension !== 'hh3' &&
                                              extension !== 'hh4'"
                                    >
                                        Comprimir y Proteger
                                    </a>
                                </li>
                                <li>
                                    <a
                                            @click="doDeHH"
                                            v-if="
                                              extension === 'hh1' ||
                                              extension === 'hh2' ||
                                              extension === 'hh3' ||
                                              extension === 'hh4'"
                                    >
                                        Desproteger y Descomprimir
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
    //import HammingOptions from "./components/HammingOptions";
    import DeHammingCorrected from "./components/DeHammingCorrected";
    import DeHammingErrors from "./components/DeHammingUncorrected";
    import AddErrors from "./components/AddErrors";
    import DeHuffman from "./components/DeHuffman";
    import Huffman from "./components/Huffman";

    export default {
        name: "app",
        components: {DeHammingCorrected, /*HammingOptions,*/ DeHammingErrors, AddErrors, DeHuffman, Huffman},
        props: {
            max: {
                type: Number,
                default: 1
            },
            value: Array
        },
        data() {
            return{
                files: [],
                file: null,
                extension: null,
                input: null,
                hamming: false,
                deHammingC: false,
                deHammingE: false,
                errors: false,
                huffman: false,
                deHuffman: false,
                message: null,
                nigth: false,
            };
        },
        mounted() {
            // Find input
            this.input = this.$el.querySelector('input[type=file]');
            this.input.addEventListener('change', () => this.onFileSelection());
            this.input.style.display = 'none';
            // Set multiple attribute on input, if max is more than one
            if (this.max > 1) {
                this.input.setAttribute('multiple', 'multiple');
            } else {
                this.input.removeAttribute('multiple');
            }
            // Populate internal value, if external value is given,
            // attempt to populate external value by firing event if not
            if (this.value) {
                this.files = this.value;
                this.file = this.value.name;
                this.extension = this.file.split('.').pop();
            } else {
                this.$emit('input', [])
            }
        },
        methods: {
            onFileSelection() {
                this.files.pop();
                // add all selected files
                for (let file of this.input.files) {
                    this.files.push({ name: file.name });
                    this.file = this.files[0].name;
                    this.extension = this.file.split('.').pop();
                }
                // reset file input
                this.input.value = null;
            },
            showUpload() {
                const event = new MouseEvent('click', {
                    'view': window,
                    'bubbles': true,
                    'cancelable': true
                });
                this.input.dispatchEvent(event)
            },
            doHamming: function () {
                 this.message = "";
                 this.hamming = true;
                 this.deHammingC = false;
                 this.deHammingE = false;
                 this.errors = false;
                 this.huffman = false;
                 this.deHuffman = false;
            },
            doHamming7(){
                window.backend.preHamming(7, "test.txt")
            },
            doDeHammingC: function () {
                this.hamming = false;
                this.message = "El archivo esta siendo corregido y decodificado";
                this.deHammingC = true;
                this.deHammingE = false;
                this.errors = false;
                this.huffman = false;
                this.deHuffman = false;
            },
            doDeHammingE: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.message = "El archivo esta siendo decodificado sin corregir errores";
                this.deHammingE = true;
                this.errors = false;
                this.huffman = false;
                this.deHuffman = false;
            },
            doErrors: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.message = "Se estan agregando errores aleatorios al archivo";
                this.errors = true;
                this.huffman = false;
                this.deHuffman = false;
            },
            doHuffman: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.errors = false;
                this.message = "El archivo esta siendo comprimido";
                this.huffman = true;
                this.deHuffman = false;
            },
            doDeHuffman: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.errors = false;
                this.huffman = false;
                this.message = "El archivo esta siendo descomprimido";
                this.deHuffman = true;
            },
            doHH: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.errors = false;
                this.huffman = false;
                this.message = "El archivo esta siendo comprimido y protegido";
                this.deHuffman = true;
            },
            doDeHH: function () {
                this.hamming = false;
                this.deHammingC = false;
                this.deHammingE = false;
                this.errors = false;
                this.huffman = false;
                this.message = "El archivo esta siendo desprotegido y descomprimido";
                this.deHuffman = true;
            },
        },

        watch: {
            files(files) {
                this.$emit('input', files)
            }
        }
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
        width: 30%;
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

    .huffman a:hover, .hamming a:hover, .hamminghuffman a:hover {
        border: 2px solid #00ff00;
        background-color: #00bb00;
    }

    .huffman a:active, .hamming a:active, .hamminghuffman a:active {
        border: 2px solid #000000;
        background-color: #00ff00;
    }

    /* Hamming Huffman */
    .hamminghuffman a {
        display: flex;
        justify-content: center;
        align-items: center;
        border: 2px solid #000000;
        border-radius: 8px;
        padding: 4px;
        width: 100%;
        height: 100%;
        background-color: #E6DE55
    }

    ul.hamminghuffman{
        list-style-type: none;
        width: 100%;
    }

    .hamminghuffman li{
        text-align: center;
        padding: 8px;
        margin: 5px;
    }

    .right {
        width: 70%;
        display: none;
    }

    .used {
        display: inherit;
        text-align: center;
    }

    .nightMode {
        background-color: #134413;
    }

</style>