<template>
  <div
    class="container d-flex align-items-center justify-content-center vh-100"
  >
    <div
      class="col-4 bg-transparent bg-opacity-25 py-5 rounded-4 border border-dark shadow-lg"
    >
      <div class="text-center mb-5">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="20vw"
          height="20vh"
          fill="currentColor"
          class="bi bi-person-circle"
          viewBox="0 0 16 16"
        >
          <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0" />
          <path
            fill-rule="evenodd"
            d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1"
          />
        </svg>
      </div>
      <BForm @submit="onSubmit"
      class="px-5 ">
        <div class="mb-5">
          <BFormGroup id="login-group-1">


              <BFormInput
                id="input-1"
                class="w-100 border-0 border-bottom border-black rounded-0 bg-transparent"
                v-model="login.email"
                type="email"
                placeholder="Email"
                required
              />

          </BFormGroup>
        </div>

        <div class="mb-5">
          <BFormGroup id="input-group-2">
            <BFormInput
              id="input-2"
              class="w-100 border-0 border-bottom border-black rounded-0 bg-transparent"
              v-model="login.pass"
              type="password"
              placeholder="Password"
              required
            />
          </BFormGroup>
        </div>
        <div class="mb-3 d-flex justify-content-between">
          <BFormGroup id="input-group-3" class="text-dark fw-light">
            <BFormCheckbox
              id="input-3"
              v-model="login.isRemember"
              class="bg-dark bg-opacity-25 border-black text-dark"
            >
              Запомнить меня
            </BFormCheckbox>
          </BFormGroup>
        </div>
        <div class="mt-5">
          <BFormGroup id="input-group-4">
            <BButton type="submit" variant="dark w-100 fs-4 fw-light"
              >Вход</BButton
            >
          </BFormGroup>
        </div>
      </BForm>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
const onSubmit = () => {
  alert(JSON.stringify(login));
};

const login = reactive({
  email: "",
  pass: "",
  isRemember: false,
});
</script>

<style>
.container {
  min-height: 100vh;
}
</style>
