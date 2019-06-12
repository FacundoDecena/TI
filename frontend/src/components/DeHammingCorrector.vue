<template>
    <div>
        <table>
            <tr>
                <td>
                    <ul>
                        <li>
                            <div class="loader"></div>
                        </li>
                        <li>
                            El archivo esta siendo corregido y decodificado
                        </li>
                    </ul>
                </td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        name: "DeHammingCorrector"
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
        background-color: #A3EB6A;
    }

    a:hover {
        border: 2px solid #00ff00;
        background-color: #00bb00;
    }

    a:active {
        border: 2px solid #000000;
        background-color: #00ff00;
    }

    .loader {
        border: 16px solid #f3f3f3;
        border-top: 16px solid #6CCA51;
        border-radius: 50%;
        width: 120px;
        height: 120px;
        animation: spin 2s linear infinite;
        margin: 0 auto;
    }

    @keyframes spin {
        0% { transform: rotate(0deg); }
        100% { transform: rotate(360deg); }
    }

</style>