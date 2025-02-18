<template>
    <div>
        Procesando autenticación con Google...
    </div>
</template>

<script setup>
import {useRouter } from "vue-router";
import axios from "axios";
import { useStore } from "../../../stores/store.js";

const router = useRouter();
const store = useStore();
const url = import.meta.env.VITE_APP_API_URL;
const code = new URLSearchParams(window.location.search).get("code");
console.log("console de code")
console.log(code)

if(code){
    authGoogle();
}else{
    console.log("No hay código")
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
    } catch (error) {
        console.log("ERRRORRR")
        console.log(error)
    }
  };

  const processAuthGoogle = async () => {
    const get_result = await goLogin();
    if (get_result) {
        console.log("get_result")
        console.log(get_result.data.token)
        store.$patch({ 
            token: get_result.data.token
        });
        router.push({ path: "/success" });
    }
  };
  processAuthGoogle();

}


</script>


