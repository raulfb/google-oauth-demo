<template>
    <div class="flex flex-col items-center justify-center">
        <h1>Successful</h1>
    </div>
    <div class="flex flex-col items-center justify-center">
        <Button @click="volver">volver</Button>
    </div>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();

function volver(){
    router.push({ path: "/" });
}


</script>