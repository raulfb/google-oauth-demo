<template>
    <div class="flex flex-col items-center justify-center" v-if="!error">
        Procesando autenticación con Google...
    </div>
    <div class="flex flex-col items-center justify-center" v-else>
       <p>Error:{{ error }}</p>
       <Button @click="volver">volver</Button>
    </div>
</template>

<script setup>
import { ref } from "vue";
import {useRouter} from "vue-router";
import axios from "axios";
import { useStore } from "../../../stores/store.js";

const router = useRouter();
const store = useStore();
const url = import.meta.env.VITE_APP_API_URL;
const code = new URLSearchParams(window.location.search).get("code");
let error = ref(false);

if(code){
    authGoogle();
}else{
    error.value = "No se ha encontrado el código de autenticación";
}

function volver(){
  router.push({ path: "/" });
}

function authGoogle() {
    const goLogin = async () => {
    let config = {
      headers: {
        accept: "application/json",
        "Content-Type": "application/json",
      },
    };

    try {
      return await axios.post(
        url + "auth/google/callback",

        {
          code: code,
        },

        config
      );
    } catch (err) {
        error.value = err.response.data.error;
    }
  };

  const processAuthGoogle = async () => {
    const get_result = await goLogin();
    if (get_result) {
        store.$patch({ 
            token: get_result.data.token
        });
        router.push({ path: "/success" });
    }
  };
  processAuthGoogle();

}
</script>


