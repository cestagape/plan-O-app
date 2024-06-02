import { createRouter, createWebHashHistory } from "vue-router";
import ClientsPage from "./pages/crm/clients/ClientsPage.vue";
import ProductsPage from "./pages/crm/products/ProductsPage.vue";
import OrdersPage from "./pages/crm/orders/OrdersPage.vue";
import LoginPage from "./pages/LoginPage.vue";
import SuppliersPage from "./pages/crm/suppliers/SuppliersPage.vue";
import LeedPage from "./pages/workflow/LeedPage/LeedPage.vue";
import CalendarPage from "./pages/workflow/CalendarPage/CalendarPage.vue";
import TaskPage from "./pages/workflow/TaskPage/TaskPage.vue";
import DashboardPage from "./pages/DashboardPage.vue";

export default createRouter ({
    history: createWebHashHistory(),
    routes: [
        { path:'/login', component: LoginPage},
        { path:'/crm/clients', component: ClientsPage},
        { path:'/crm/products', component: ProductsPage},
        { path:'/crm/orders', component: OrdersPage},
        { path:'/crm/suppliers', component: SuppliersPage},
        { path:'/workflow/leeds', component: LeedPage},
        { path:'/workflow/tasks', component: TaskPage},
        { path:'/workflow/calendar', component: CalendarPage},
        { path:'/dashboard', component: DashboardPage}
    ]
})