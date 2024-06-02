import { createRouter, createWebHashHistory } from "vue-router";
import ClientsPage from "./pages/clients/ClientsPage.vue";
import ProductsPage from "./pages/products/ProductsPage.vue";
import OrdersPage from "./pages/orders/OrdersPage.vue";

export default createRouter ({
    history: createWebHashHistory(),
    routes: [
        { path:'/crm/clients', component: ClientsPage},
        { path:'/crm/products', component: ProductsPage},
        { path:'/crm/orders', component: OrdersPage},
        { path:'/', redirect: '/crm/clients'},

    ]
})