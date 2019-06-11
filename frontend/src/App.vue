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
                    <tr class="opciones">
                        <td>
                            <ul class="hamming">
                                <li>
                                    <a>
                                        Codificar con Hamming
                                    </a>
                                </li>
                                <li class="decodificar">
                                    <a>
                                        Decodificar Hamming
                                    </a>
                                </li>
                                <li class="aclaracion">
                                    Corrigiendo errores
                                </li>
                                <li class="decodificar">
                                    <a>
                                        Decodificar Hamming
                                    </a>
                                </li>
                                <li class="aclaracion">
                                    Sin corregir errores
                                </li>
                                <li>
                                    <a>
                                        Introducir errores
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
    export default {
        name: "app",
        data() {

        },
        methods: {
            loadFile: function() {
                window.backend.Todos.LoadNewList();
            },
        },
    };
</script>
<style>
    * {
        box-sizing: border-box;
    }
    /* html */
    html{
        text-align: center;
        margin-top: 20px;
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
        width: 50%;
        min-width: max-content;
    }
    /* select action */
    .hamming a {
        display: flex;
        justify-content: center;
        align-items: center;
        border: 2px solid #000000;
        border-radius: 8px;
        padding: 4px;
        width: 100%;
        height: 100%;
        background-color: green;
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
    }


</style>