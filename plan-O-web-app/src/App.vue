<template>
  <div class="container-fluid">
    <div class="row flex-nowrap">
      <Sidebar
        class="col-auto col-xxl-2 col-xl-2 col-md-3 col-sm-3 shadow-lg"
      />
      <!-- <router-view></router-view> -->
      <KanbanBoard> 
      </KanbanBoard>
    </div>
  </div>
</template>

<script>
import Sidebar from "./modules/Sidebar.vue";
import ClientsPage from "./pages/clients/ClientsPage.vue";
import ProductsPage from "./pages/products/ProductsPage.vue";
import OrdersPage from "./pages/orders/OrdersPage.vue";
import KanbanBoard from './modules/Kanban/KanbanBoard.vue'
import KanbanColumn from './modules/Kanban/KanbanColumn.vue'



export default {
  components: {
    ClientsPage,
    ProductsPage,
    Sidebar,
    OrdersPage
  },
  data() {
    return {};
  },
};
</script>
<style>
body {
  height: 100%;
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
</style>
