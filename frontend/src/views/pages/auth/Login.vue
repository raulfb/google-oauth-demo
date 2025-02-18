<template>
  <div
    class="h-screen text-slate-200"
    style="background: linear-gradient(180deg, #525d72 0%, #69768d 100%)"
  >
    <div class="flex justify-center items-center h-full">
      <form @submit.prevent="login()" class="justify-center text-center">
        <span class="block text-white text-3xl mb-4">
          Iniciar sesión
        </span>

        <Button
          label="Iniciar sesión con Google"
          @click="loginConGoogle"
          class="w-12/12 mb-6 p-button-outlined p-button-google"
        >
          <i class="pi pi-google"></i> Iniciar sesión con Google
        </Button>
        <Message
          class="mt-5 ml-5 mr-5"
          v-show="loginAlert"
          id="mensaje"
          severity="error"
          :closable="false"
        >
          El email o la contraseña son incorrectos. Por favor inténtelo de nuevo.
        </Message>
      </form>
    </div>
    <Toast position="bottom-right" />
  </div>
</template>
<script setup>
import { ref } from "vue";
import { useToast } from "primevue/usetoast";

import axios from "axios";

const url = import.meta.env.VITE_APP_API_URL;
let loginAlert = ref(false);
const toast = useToast();

function loginConGoogle() {
  const goLoginGoogle = async () => {
    let config = {
      headers: {
        accept: "application/json",
        "Content-Type": "application/json",
      },
    };

    try {
      return await axios.get(
        url + "auth/google/url",

        config
      );
    } catch (error) {
      toast.add({
          severity: "error",
          summary: "Error",
          detail: "Error al iniciar sesión con Google",
          life: 2500,
        });
    }
  };

  const processLoginGoogle = async () => {
    const get_result = await goLoginGoogle();
    if (get_result) {
      window.location.href = get_result.data.url
    }
  };
  processLoginGoogle();
}

</script>

<style scoped>
.p-button-google {
  background-color: transparent; /* Fondo transparente */
  color: white; /* Color del texto a blanco */
  border: 2px solid #4285f4; /* Borde del color de Google */
}

.p-button-google:hover {
  background-color: rgba(66, 133, 244, 0.1); /* Color de fondo al pasar el mouse */
}

.p-button-google i {
  margin-right: 0.5rem; /* Espacio entre el icono y el texto */
  /* Eliminar el color del icono para que use el color original de PrimeIcons */
}
</style>