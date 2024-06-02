<template>
  <div class="bg-dark px-0 dark bg-opacity-75">
    <div
      class="d-flex flex-column align-items-center align-items-sm-start ps-3 py-2 min-vh-100 max-vh-100"
      style="max-height: 100vh; min-height: 100vh"
    >
      <a
        href="/"
        class="d-flex align-items-center p-0 m-0 me-md-auto text-white text-decoration-none"

      >
        <span class="fs-5 d-none d-sm-inline">Plan-Ó</span>
      </a>
      <hr />
      <div class="navbar-nav-scroll pe-2 w-100 menu" style="">
        <p
          class="fw-lighter fs-6 text-white-50 mx-0"
          v-if="window.width >= 576"
        >
          Menu
        </p>
        <ul
          class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start text-white"
          id="menu"
        >
          <li>
            <a href="#" class="nav-link px-0 align-middle text-reset d-flex">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="currentColor"
                transform="scale(1,-1)"
                class="bi bi-kanban-fill nav-svg"
                viewBox="0 0 16 16"
              >
                <path
                  d="M2.5 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm5 2h1a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1m-5 1a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1zm9-1h1a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1"
                />
              </svg>
              <span class="ms-3 d-none d-sm-inline">Dashboard</span>
            </a>
          </li>
          <li>
            <a href="#" class="nav-link px-0 align-middle text-reset d-flex">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="currentColor"
                class="bi bi-envelope"
                viewBox="0 0 16 16"
              >
                <path
                  d="M0 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2zm2-1a1 1 0 0 0-1 1v.217l7 4.2 7-4.2V4a1 1 0 0 0-1-1zm13 2.383-4.708 2.825L15 11.105zm-.034 6.876-5.64-3.471L8 9.583l-1.326-.795-5.64 3.47A1 1 0 0 0 2 13h12a1 1 0 0 0 .966-.741M1 11.105l4.708-2.897L1 5.383z"
                />
              </svg>
              <span class="ms-3 d-none d-sm-inline">Почта</span>
            </a>
          </li>
          <li>
            <BCollapse id="collapse-1">
              <template #header="{ visible, toggle, id }">
                <a
                  style="cursor: pointer;"
                  :aria-expanded="visible"
                  :aria-controls="id"
                  @click="toggle"
                  class="nav-link px-0 align-middle d-flex"
                  :class="visible ? 'text-info' : 'text-white'"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    fill="currentColor"
                    class="bi bi-display"
                    viewBox="0 0 16 16"
                  >
                    <path
                      d="M0 4s0-2 2-2h12s2 0 2 2v6s0 2-2 2h-4q0 1 .25 1.5H11a.5.5 0 0 1 0 1H5a.5.5 0 0 1 0-1h.75Q6 13 6 12H2s-2 0-2-2zm1.398-.855a.76.76 0 0 0-.254.302A1.5 1.5 0 0 0 1 4.01V10c0 .325.078.502.145.602q.105.156.302.254a1.5 1.5 0 0 0 .538.143L2.01 11H14c.325 0 .502-.078.602-.145a.76.76 0 0 0 .254-.302 1.5 1.5 0 0 0 .143-.538L15 9.99V4c0-.325-.078-.502-.145-.602a.76.76 0 0 0-.302-.254A1.5 1.5 0 0 0 13.99 3H2c-.325 0-.502.078-.602.145"
                    />
                  </svg>
                  <span class="ms-3 d-none d-sm-inline">CRM</span>
                </a>
              </template>
              <ul
                class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start text-white ms-0"
                v-if="window.width >= 576"
              >
                <div class="ms-3">
                  <li style="margin-left: 23px">
                    <a
                      href="#/crm/orders"
                      class="nav-link px-0 align-middle text-reset d-flex"
                      ><span class="d-none d-sm-inline">Заказы</span></a
                    >
                  </li>
                  <li style="margin-left: 23px">
                    <a
                      href="#/crm/products"
                      class="nav-link px-0 align-middle text-reset d-flex"
                      ><span class="d-none d-sm-inline">Товары</span></a
                    >
                  </li>
                  <li style="margin-left: 23px">
                    <a
                      href="#"
                      class="nav-link px-0 align-middle text-reset d-flex"
                      ><span class="d-none d-sm-inline">Поставщики</span></a
                    >
                  </li>
                  <li style="margin-left: 23px">
                    <a
                      href="#/crm/clients"
                      class="nav-link px-0 align-middle text-reset d-flex"
                      ><span class="d-none d-sm-inline">Клиенты</span></a
                    >
                  </li>
                </div>
              </ul>
            </BCollapse>
          </li>
          <li>
            <BCollapse id="collapse-2">
              <template #header="{ visible, toggle, id }">
                <a
                  style="cursor: pointer;"
                  variant="primary"
                  :aria-expanded="visible"
                  :aria-controls="id"
                  @click="toggle"
                  class="nav-link px-0 align-middle d-flex"
                  :class="visible ? 'text-info' : 'text-white'"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    fill="currentColor"
                    class="bi bi-graph-up-arrow"
                    viewBox="0 0 16 16"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M0 0h1v15h15v1H0zm10 3.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-1 0V4.9l-3.613 4.417a.5.5 0 0 1-.74.037L7.06 6.767l-3.656 5.027a.5.5 0 0 1-.808-.588l4-5.5a.5.5 0 0 1 .758-.06l2.609 2.61L13.445 4H10.5a.5.5 0 0 1-.5-.5"
                    />
                  </svg>
                  <span class="ms-3 d-none d-sm-inline">Workflow</span>
                </a>
              </template>
              <ul
                class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start text-white ms-0"
                v-if="window.width >= 576"
              >
                <div class="ms-3">
                  <li style="margin-left: 23px">
                    <a
                      href="#"
                      class="nav-link px-0 align-middle text-reset d-flex"
                    >
                      <span class="d-none d-sm-inline">Задачи</span>
                    </a>
                  </li>
                  <li style="margin-left: 23px">
                    <a
                      href="#"
                      class="nav-link px-0 align-middle text-reset d-flex"
                      ><span class="d-none d-sm-inline">Лиды</span></a
                    >
                  </li>
                  <li style="margin-left: 23px">
                    <a
                      href="#"
                      class="nav-link px-0 align-middle text-reset d-flex"
                      ><span class="d-none d-sm-inline">Календарь</span></a
                    >
                  </li>
                </div>
              </ul>
            </BCollapse>
          </li>
        </ul>
      </div>

      <hr />
      <ul
        class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start text-white"
        id="menu"
      >
        <li>
          <a href="#" class="nav-link px-0 align-middle text-reset d-flex">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor"
              class="bi bi-question-square"
              viewBox="0 0 16 16"
            >
              <path
                d="M14 1a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1zM2 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2z"
              />
              <path
                d="M5.255 5.786a.237.237 0 0 0 .241.247h.825c.138 0 .248-.113.266-.25.09-.656.54-1.134 1.342-1.134.686 0 1.314.343 1.314 1.168 0 .635-.374.927-.965 1.371-.673.489-1.206 1.06-1.168 1.987l.003.217a.25.25 0 0 0 .25.246h.811a.25.25 0 0 0 .25-.25v-.105c0-.718.273-.927 1.01-1.486.609-.463 1.244-.977 1.244-2.056 0-1.511-1.276-2.241-2.673-2.241-1.267 0-2.655.59-2.75 2.286m1.557 5.763c0 .533.425.927 1.01.927.609 0 1.028-.394 1.028-.927 0-.552-.42-.94-1.029-.94-.584 0-1.009.388-1.009.94"
              />
            </svg>
            <span class="ms-3 d-none d-sm-inline">Помощь</span>
          </a>
        </li>
        <li>
          <a href="#" class="nav-link px-0 align-middle text-reset d-flex">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor"
              class="bi bi-gear"
              viewBox="0 0 16 16"
            >
              <path
                d="M8 4.754a3.246 3.246 0 1 0 0 6.492 3.246 3.246 0 0 0 0-6.492M5.754 8a2.246 2.246 0 1 1 4.492 0 2.246 2.246 0 0 1-4.492 0"
              />
              <path
                d="M9.796 1.343c-.527-1.79-3.065-1.79-3.592 0l-.094.319a.873.873 0 0 1-1.255.52l-.292-.16c-1.64-.892-3.433.902-2.54 2.541l.159.292a.873.873 0 0 1-.52 1.255l-.319.094c-1.79.527-1.79 3.065 0 3.592l.319.094a.873.873 0 0 1 .52 1.255l-.16.292c-.892 1.64.901 3.434 2.541 2.54l.292-.159a.873.873 0 0 1 1.255.52l.094.319c.527 1.79 3.065 1.79 3.592 0l.094-.319a.873.873 0 0 1 1.255-.52l.292.16c1.64.893 3.434-.902 2.54-2.541l-.159-.292a.873.873 0 0 1 .52-1.255l.319-.094c1.79-.527 1.79-3.065 0-3.592l-.319-.094a.873.873 0 0 1-.52-1.255l.16-.292c.893-1.64-.902-3.433-2.541-2.54l-.292.159a.873.873 0 0 1-1.255-.52zm-2.633.283c.246-.835 1.428-.835 1.674 0l.094.319a1.873 1.873 0 0 0 2.693 1.115l.291-.16c.764-.415 1.6.42 1.184 1.185l-.159.292a1.873 1.873 0 0 0 1.116 2.692l.318.094c.835.246.835 1.428 0 1.674l-.319.094a1.873 1.873 0 0 0-1.115 2.693l.16.291c.415.764-.42 1.6-1.185 1.184l-.291-.159a1.873 1.873 0 0 0-2.693 1.116l-.094.318c-.246.835-1.428.835-1.674 0l-.094-.319a1.873 1.873 0 0 0-2.692-1.115l-.292.16c-.764.415-1.6-.42-1.184-1.185l.159-.291A1.873 1.873 0 0 0 1.945 8.93l-.319-.094c-.835-.246-.835-1.428 0-1.674l.319-.094A1.873 1.873 0 0 0 3.06 4.377l-.16-.292c-.415-.764.42-1.6 1.185-1.184l.292.159a1.873 1.873 0 0 0 2.692-1.115z"
              />
            </svg>
            <span class="ms-3 d-none d-sm-inline">Настройки</span>
          </a>
        </li>
      </ul>
      <hr />
      <ul
        class="nav nav-pills flex-column align-items-sm-end text-white py-0 my-0 ms-0"
        id="menu"
      >
        <li>
          <a
          class="nav-link px-0 text-reset d-flex ms-0">
            <BDropdown
              dropend
              variant="link text-reset"
              text=""
              toggle-class="text-decoration-none"
              class="ms-0"
              v-bind:no-caret="window.width >= 576 ? False : ''"
            >
              <template #button-content>
                <img
                  src="https://avatars.githubusercontent.com/u/67867944?v=4"
                  width="24"
                  height="24"
                  class="rounded-circle"
                  :class="
                    window.width < 992 && window.width >= 576 ? 'me-3' : 'me-0'
                  "
                />
                <span
                  class="ms-3 d-none d-sm-inline d-md-inline"
                  v-if="window.width >= 992"
                  >{{ profile.name }} {{ profile.surname }}</span
                >
              </template>
              <BDropdownItem href="#" class="text-reset fw-bolder" variant="dark">
                Выход
              </BDropdownItem>
            </BDropdown>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      profile: {
        name: "Имя",
        surname: "Фамилия",
      },
      window: {
        width: 0,
        height: 0,
      },
      showProfileName: true,
    };
  },
  created() {
    window.addEventListener("resize", this.handleResize);
    this.handleResize();
  },
  destroyed() {
    window.removeEventListener("resize", this.handleResize);
  },
  methods: {
    handleResize() {
      this.window.width = window.innerWidth;
      this.window.height = window.innerHeight;
    },
  },
};
</script>

<style>
a:hover {
  color: #6c757d !important;
}
.nav {
  margin-left: 5%;
}
hr {
  width: 100%;
  color: white;
}
.menu {
  overflow-y: auto;
  overflow-x: hidden;
  scrollbar-width: thin;
  scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
  scroll-margin: 0;
  scroll-padding: 0;
  scrollbar-gutter: 0;
}
.text-collapsed {
  color: gray; /* Color when collapsed */
}
.text-expanded {
  color: black; /* Color when expanded */
}
</style>
