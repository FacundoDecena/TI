<template>
    <div>
        <table>
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
</template>

<script>
    //import App from "../App"
    export default {
        name: "HammingOptions",
        methods: {
            doHamming7(){
                window.backend.preHamming(7, "test.txt")
            }
        }
    }
</script>

<style scoped>
    ul {
        list-style-type: none;
        width: 100%;
    }
    li{
        text-align: center;
        padding: 8px;
        margin: 0;
    }
    a {
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

    a:hover {
        border: 2px solid #00ff00;
        background-color: #00bb00;
    }

    a:active {
        border: 2px solid #000000;
        background-color: #00ff00;
    }

</style>