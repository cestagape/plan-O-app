<template>
  <div class="container-fluid">
    <div v-if="!hideSidebarAndNavbar" class="row flex-nowrap">
      <Sidebar class="col-auto col-xxl-2 col-xl-2 col-md-3 col-sm-3 shadow-lg" />
      <div class="content col-auto col-xxl-10 col-xl-10 col-md-9 col-sm-9 align-items-start mx-0 px-0 mh-100">
        <BNavbar
          toggleable="expand-lg"
          variant="dark bg-opacity-75"
          v-b-color-mode="'dark'"
          class="align-self-start align-items-center shadow-lg"
        >
          <BNavbarNav>
            <BBreadcrumb class="text-decoration-none" />
          </BNavbarNav>
          <BNavbarNav class="w-50">
            <BNavForm class="d-flex w-100">
              <BFormInput id="input-1" class="mx-2" placeholder="Поиск..." />
              <BButton type="submit" variant="outline-light">Искать</BButton>
            </BNavForm>
          </BNavbarNav>
        </BNavbar>
        <router-view></router-view>
      </div>
    </div>
    <div v-else>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Sidebar from "./modules/Sidebar.vue";

export default {
  components: {
    Sidebar,
  },
  computed: {
    hideSidebarAndNavbar() {
      return this.$route.meta.hideSidebarAndNavbar;
    }
  }
};
</script>

<style>
body {
  background: radial-gradient(
    circle,
    #ec8f5e,
    #f3b664,
    #f1eb8b,
    #cee6a9,
    #a0bd73,
    #83b58b,
    #8cc0bc,
    #8cc0de,
    #2c4d94,
    #1c315e
  );
  background-size: 400% 400%;
  animation: gradient 40s ease infinite;
}
@keyframes gradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}
.content {
  max-height: 100vh;
  min-height: 100vh;
}
</style>
