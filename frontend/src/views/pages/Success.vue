<template>
    <div class="flex flex-col items-center justify-center">
        <p>{{ mensaje }}</p>
    </div>
    <div class="flex flex-col items-center justify-center">
        <Button @click="volver">volver</Button>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "../../../src/stores/store.js";
import axios from "axios";

const router = useRouter();
const store = useStore();
const url = import.meta.env.VITE_APP_API_URL;
let mensaje = ref("");

function volver(){
    router.push({ path: "/" });
}

function getRutaProtegida() {
  const goRutaProtegida = async () => {
    let config = {
      headers: {
        accept: "application/json",
        "Content-Type": "application/json",
        "Authorization": store.token,
      },
    };

    try {
      return await axios.get(
        url + "protegida",
        config
      );
    } catch (error) {
        mensaje.value = "Error al obtener la ruta protegida";
    }
  };

  const processRutaProtegida = async () => {
    const get_result = await goRutaProtegida();
    if (get_result) {
        mensaje.value = get_result.data.message;
    }
  };
  processRutaProtegida();
}

getRutaProtegida();
</script>