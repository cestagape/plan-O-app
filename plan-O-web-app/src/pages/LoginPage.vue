<template>
  <div
    class="container d-flex align-items-center justify-content-center vh-100"
  >
    <div
      class="col-4 bg-transparent bg-opacity-25 px-5 py-5 rounded-4 border border-dark shadow-lg"
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
      <BForm @submit="onSubmit">
        <div class="mb-5">
          <BFormGroup id="login-group-1">

              <!-- <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-envelope-fill"
                viewBox="0 0 16 16"
              >
                <path
                  d="M.05 3.555A2 2 0 0 1 2 2h12a2 2 0 0 1 1.95 1.555L8 8.414zM0 4.697v7.104l5.803-3.558zM6.761 8.83l-6.57 4.027A2 2 0 0 0 2 14h12a2 2 0 0 0 1.808-1.144l-6.57-4.027L8 9.586zm3.436-.586L16 11.801V4.697z"
                />
              </svg> -->
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
            <!-- <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              class="bi bi-lock-fill"
              viewBox="0 0 16 16"
            >
              <path
                d="M8 1a2 2 0 0 1 2 2v4H6V3a2 2 0 0 1 2-2m3 6V3a3 3 0 0 0-6 0v4a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2"
              />
            </svg> -->
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
