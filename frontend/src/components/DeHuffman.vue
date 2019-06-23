<template>
    <div class="loader"></div>
</template>

<script>
    export default {
        name: "deHuffman"
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

    .loader {
        border: 16px solid #f3f3f3;
        border-top: 16px solid #A3EB6A;
        border-radius: 50%;
        width: 120px;
        height: 120px;
        animation: spin 2s linear infinite;
        margin-left: 350px;
    }

    @keyframes spin {
        0% { transform: rotate(0deg); }
        100% { transform: rotate(360deg); }
    }

</style>