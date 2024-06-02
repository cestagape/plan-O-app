<template>
  <BContainer class="py-2 px-2 mx-0" :fluid="true">
    <BButtonGroup class="mb-2">
      <BButton
        :variant="kanbanShow ? 'dark' : 'outline-light'"
        @click="
          () => {
            kanbanShow = true;
            tableShow = false;
          }
        "
        >Канбан</BButton
      >
      <BButton
        :variant="tableShow ? 'dark' : 'outline-light'"
        @click="
          () => {
            kanbanShow = false;
            tableShow = true;
          }
        "
        >Таблица</BButton
      >
    </BButtonGroup>

    <BCardGroup deck v-if="kanbanShow">
      <BCard
        header="В обработке"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 broder-0"
        bodyClass="p-2 cardin text-center"
      >
        <BButton
          pill
          variant="outline-dark"
          class="mb-2"
          @click="add(inProcess)"
        >
          Добавить</BButton
        >
        <draggable
          class="list-group"
          :list="inProcess"
          group="people"
          @change="log"
          itemKey="name"
        >
          <template #item="{ element, index }">
            <div class="list-group-item">{{ element.name }} {{ index }}</div>
          </template>
        </draggable>
      </BCard>
      <BCard
        header="Ожидает оплаты"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 broder-0"
        bodyClass="p-2 cardin text-center"
      >
        <BButton
          pill
          variant="outline-dark"
          class="mb-2"
          @click="add(paymentPending)"
        >
          Добавить</BButton
        >
        <draggable
          class="list-group"
          :list="paymentPending"
          group="people"
          @change="log"
          itemKey="name"
        >
          <template #item="{ element, index }">
            <div class="list-group-item">{{ element.name }} {{ index }}</div>
          </template>
        </draggable>
      </BCard>
      <BCard
        header="В производстве"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 broder-0"
        bodyClass="p-2 cardin text-center"
      >
        <BButton
          pill
          variant="outline-dark"
          class="mb-2"
          @click="add(inProduction)"
        >
          Добавить</BButton
        >
        <draggable
          class="list-group"
          :list="inProduction"
          group="people"
          @change="log"
          itemKey="name"
        >
          <template #item="{ element, index }">
            <div class="list-group-item">{{ element.name }} {{ index }}</div>
          </template>
        </draggable>
      </BCard>
      <BCard
        header="Контроль качества"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 broder-0"
        bodyClass="p-2 cardin text-center"
      >
        <BButton
          pill
          variant="outline-dark"
          class="mb-2"
          @click="add(qualityControl)"
        >
          Добавить</BButton
        >
        <draggable
          class="list-group"
          :list="qualityControl"
          group="people"
          @change="log"
          itemKey="name"
        >
          <template #item="{ element, index }">
            <div class="list-group-item">{{ element.name }} {{ index }}</div>
          </template>
        </draggable>
      </BCard>
      <BCard
        header="Упаковка"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 broder-0"
        bodyClass="p-2 cardin text-center"
      >
        <BButton
          pill
          variant="outline-dark"
          class="mb-2"
          @click="add(packaging)"
        >
          Добавить</BButton
        >
        <draggable
          class="list-group"
          :list="packaging"
          group="people"
          @change="log"
          itemKey="name"
        >
          <template #item="{ element, index }">
            <div class="list-group-item">{{ element.name }} {{ index }}</div>
          </template>
        </draggable>
      </BCard>
      <BCard
        header="Печать документов"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 broder-0"
        bodyClass="p-2 cardin text-center"
      >
        <BButton
          pill
          variant="outline-dark"
          class="mb-2"
          @click="add(docsCreation)"
        >
          Добавить</BButton
        >
        <draggable
          class="list-group"
          :list="docsCreation"
          group="people"
          @change="log"
          itemKey="name"
        >
          <template #item="{ element, index }">
            <div class="list-group-item">{{ element.name }} {{ index }}</div>
          </template>
        </draggable>
      </BCard>
      <BCard
        header="Передано в доставку"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 broder-0"
        bodyClass="p-2 cardin text-center"
      >
        <BButton
          pill
          variant="outline-dark"
          class="mb-2"
          @click="add(sentForDelivery)"
        >
          Добавить</BButton
        >
        <draggable
          class="list-group"
          :list="sentForDelivery"
          group="people"
          @change="log"
          itemKey="name"
        >
          <template #item="{ element, index }">
            <div class="list-group-item">{{ element.name }} {{ index }}</div>
          </template>
        </draggable>
      </BCard>
    </BCardGroup>
    <div v-if="tableShow">
      <div class="d-flex mb-2" v-if="tableShow">
        <BButton
          type="submit"
          class="text-info"
          variant="dark me-2"
          @click="add"
        >
          Добавить</BButton
        >
        <BInputGroup class="w-50 d-flex">
          <BFormInput
            id="input-2"
            placeholder="Введите имя"
            v-model="filter"
            type="search"
          />
          <BInputGroupAppend>
            <BButton :disabled="!filter" @click="filter = ''">Clear</BButton>
          </BInputGroupAppend>
        </BInputGroup>
      </div>
      <BTable
        class="mt-1 rounded-3 order-table"
        style="max-height: 75vh; max-width: 100%"
        model="sortBy"
        :sort-internal="true"
        :items="itemsTyped"
        :fields="fieldsTyped"
        :filter="filter"
        :responsive="False"
        :filterable="filterOn"
        :multisort="true"
        :stickyHeader="true"
        @filtered="onFiltered"
      >
        <template #cell(actions)="row">
          <BButton
            size="sm"
            class="me-1 my-1 fw-light"
            variant="light"
            bg
            @click="info(row.item, row.index)"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              class="bi bi-eye me-1"
              viewBox="0 0 16 16"
            >
              <path
                d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8M1.173 8a13 13 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5s3.879 1.168 5.168 2.457A13 13 0 0 1 14.828 8q-.086.13-.195.288c-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5s-3.879-1.168-5.168-2.457A13 13 0 0 1 1.172 8z"
              />
              <path
                d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5M4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0"
              /></svg>
              Просмотр</BButton
          >
          <BButton
            size="sm"
            class="me-1 fw-light"
            variant="light"
            @click="edit(row.item, row.index)"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              class="bi bi-pencil-fill me-1"
              viewBox="0 0 16 16">
              <path
                d="M12.854.146a.5.5 0 0 0-.707 0L10.5 1.793 14.207 5.5l1.647-1.646a.5.5 0 0 0 0-.708zm.646 6.061L9.793 2.5 3.293 9H3.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.207zm-7.468 7.468A.5.5 0 0 1 6 13.5V13h-.5a.5.5 0 0 1-.5-.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.5-.5V10h-.5a.5.5 0 0 1-.175-.032l-.179.178a.5.5 0 0 0-.11.168l-2 5a.5.5 0 0 0 .65.65l5-2a.5.5 0 0 0 .168-.11z"/></svg>
                Изменить</BButton>
        </template>
      </BTable>
      
      <!-- Info modal -->
      <BModal
        size="lg"
        :id="infoOrder.id"
        v-model="infoOrder.open"
        :title="infoOrder.title"
        :ok-only="true"
        @hide="resetinfoOrder"
      >
        <BForm v-if="show">
          <!-- Название компании -->
          <BFormGroup
            id="info-input-group-1"
            label="Название компании"
            label-for="info-input-1"
            class="mb-1"
          >
            <BInputGroup>
              <BFormInput
                id="info-input-1"
                v-model="infoOrder.content.companyName"
                type="email"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextOrder" variant="outline-dark"
                  >Copy</BButton
                >
              </BInputGroupAppend>
            </BInputGroup>
          </BFormGroup>

          <!-- E-mail -->
          <BFormGroup
            id="info-input-group-2"
            label="Email"
            label-for="info-input-2"
            class="my-1"
          >
            <BInputGroup>
              <BFormInput
                id="info-input-2"
                v-model="infoOrder.content.email"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextOrder(info)" variant="outline-dark"
                  >Copy</BButton
                >
              </BInputGroupAppend>
            </BInputGroup>
          </BFormGroup>

          <!-- Тип клиента -->
          <BFormGroup
            id="input-group-3"
            label="Тип клиента"
            label-for="info-input-3"
            class="my-1"
          >
            <BInputGroup>
              <BFormInput
                id="info-input-3"
                v-model="infoOrder.content.customerType"
                type="text"
                disabled
              />
            </BInputGroup>
          </BFormGroup>

          <!-- Адрес -->
          <BFormGroup
            id="info-input-group-4"
            label="Адрес"
            label-for="info-input-4"
            class="my-1"
          >
            <BInputGroup>
              <BFormInput
                id="info-input-4"
                v-model="infoOrder.content.adress"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextOrder" variant="outline-dark"
                  >Copy</BButton
                >
              </BInputGroupAppend>
            </BInputGroup>
          </BFormGroup>

          <!-- Номер телефона -->
          <BFormGroup
            id="info-input-group-5"
            label="Номер телефона"
            label-for="info-input-5"
            class="my-1"
          >
            <BInputGroup>
              <BFormInput
                id="info-input-5"
                v-model="infoOrder.content.phone"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextOrder" variant="outline-dark"
                  >Copy</BButton
                >
              </BInputGroupAppend>
            </BInputGroup>
          </BFormGroup>

          <!-- Имя менеджера -->
          <BFormGroup
            id="info-input-group-6"
            label="Имя менеджера"
            label-for="info-input-6"
            class="my-1"
          >
            <BInputGroup>
              <BFormInput
                id="info-input-6"
                v-model="infoOrder.content.managerName"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextOrder" variant="outline-dark"
                  >Copy</BButton
                >
              </BInputGroupAppend>
            </BInputGroup>
          </BFormGroup>

          <!-- Комментарии -->
          <BFormGroup
            id="info-input-group-7"
            label="Комментарии"
            label-for="info-input-7"
            class="my-1"
          >
            <BInputGroup>
              <BFormInput
                id="info-input-7"
                v-model="infoOrder.content.comments"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextOrder" variant="outline-dark"
                  >Copy</BButton
                >
              </BInputGroupAppend>
            </BInputGroup>
          </BFormGroup>
        </BForm>
      </BModal>
    </div>

    <rawDisplayer class="col-2" :value="inProcess" title="List 1" />
    <rawDisplayer class="col-2" :value="paymentPending" title="List 2" />
    <rawDisplayer class="col-2" :value="paymentPending" title="List 3" />
  </BContainer>
</template>
<script setup lang="ts">
import draggable from "vuedraggable";
import {
  BButton,
  BFormSelect,
  BInputGroup,
  BFormGroup,
  BCol,
  BFormInput,
  BInputGroupAppend,
  BPagination,
  BRow,
  BModal,
  BContainer,
  BTable,
  type BTableSortBy,
  type ColorVariant,
  type TableFieldRaw,
  type TableItem,
} from "bootstrap-vue-next";
import { reactive, ref, defineComponent } from "vue";

const kanbanShow = ref(true);
const tableShow = ref(false);

interface Order {
  id: number;
  createdAt: string; // ISO 8601 date string for TIMESTAMPTZ
  design: string;
  quantity: number;
  status: string;
  price: number;
  addr: string;
  employee: number;
  deadline: string; // ISO 8601 date string for TIMESTAMPTZ
  paymentType: string;
  deliveryInfo: Record<string, any>; // JSONB field
  additional: Record<string, any>; // JSONB field
}

const statuses = [
  "inProcess",
  "paymentPending",
  "inProduction",
  "qualityControl",
  "packaging",
  "docsCreation",
  "sentForDelivery",
];

const paymentTypes = ["рассчетный счет компании", "наличные", "сбп"];

const fieldsTyped: Exclude<TableFieldRaw<Order>, string>[] = [
  {
    key: "id",
    label: "ID заказа",
    sortable: true,
    sortDirection: "asc",
  },
  {
    key: "createdAt",
    label: "Размещен",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "quantity",
    label: "Количество",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "status",
    label: "Статус",
    sortable: true,
    sortDirection: "asc",
  },
  {
    key: "price",
    label: "Price",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "employee",
    label: "Ответственный",
    sortable: true,
    sortDirection: "asc",
  },
  {
    key: "deadline",
    label: "Дэдлайн",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "paymentType",
    label: "Тип оплаты",
    sortable: true,
    sortDirection: "asc",
  },
  {
    key: "actions",
    label: "Действия",
  },
];

const infoOrder = reactive({
  open: false,
  id: "info-modal",
  title: "",
  content: {},
});

function info(item: TableItem<Order>, index: number) {
  infoOrder.title = `# ${item.id}`;
  infoOrder.content = item;
  infoOrder.open = true;
}
function resetInfoOrder() {
  infoOrder.title = "";
  infoOrder.content = "";
}
function resetEditOrder() {
  infoOrder.title = "";
  infoOrder.content = "";
}
function resetAddOrder() {
  infoOrder.title = "";
  infoOrder.content = "";
}

const itemsTyped: TableItem<Order>[] = [
  {
    id: 1,
    createdAt: new Date().toISOString(),
    design: "Design A",
    quantity: 10,
    status: statuses[0], // inProcess
    price: 1000,
    addr: "123 Main St",
    employee: 101,
    deadline: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString(), // 1 week from now
    paymentType: paymentTypes[0], // рассчетный счет компании
    deliveryInfo: { courier: "DHL", trackingNumber: "123456789" },
    additional: { notes: "Handle with care" },
  },
  {
    id: 2,
    createdAt: new Date().toISOString(),
    design: "Design B",
    quantity: 5,
    status: statuses[1], // paymentPending
    price: 500,
    addr: "456 Elm St",
    employee: 102,
    deadline: new Date(Date.now() + 14 * 24 * 60 * 60 * 1000).toISOString(), // 2 weeks from now
    paymentType: paymentTypes[1], // наличные
    deliveryInfo: { courier: "UPS", trackingNumber: "987654321" },
    additional: { notes: "Fragile item" },
  },
  {
    id: 3,
    createdAt: new Date().toISOString(),
    design: "Design C",
    quantity: 20,
    status: statuses[2], // inProduction
    price: 2000,
    addr: "789 Oak St",
    employee: 103,
    deadline: new Date(Date.now() + 3 * 24 * 60 * 60 * 1000).toISOString(), // 3 days from now
    paymentType: paymentTypes[2], // сбп
    deliveryInfo: { courier: "FedEx", trackingNumber: "1122334455" },
    additional: { notes: "Urgent delivery" },
  },
  {
    id: 4,
    createdAt: new Date().toISOString(),
    design: "Design D",
    quantity: 15,
    status: statuses[3], // qualityControl
    price: 1500,
    addr: "101 Maple St",
    employee: 104,
    deadline: new Date(Date.now() + 10 * 24 * 60 * 60 * 1000).toISOString(), // 10 days from now
    paymentType: paymentTypes[0], // рассчетный счет компании
    deliveryInfo: { courier: "DHL", trackingNumber: "2233445566" },
    additional: { notes: "Inspect carefully" },
  },
  {
    id: 5,
    createdAt: new Date().toISOString(),
    design: "Design E",
    quantity: 25,
    status: statuses[4], // packaging
    price: 2500,
    addr: "202 Birch St",
    employee: 105,
    deadline: new Date(Date.now() + 5 * 24 * 60 * 60 * 1000).toISOString(), // 5 days from now
    paymentType: paymentTypes[1], // наличные
    deliveryInfo: { courier: "UPS", trackingNumber: "5566778899" },
    additional: { notes: "Package securely" },
  },
];

const totalRows = ref(itemsTyped.length);
const currentPage = ref(1);
const perPage = ref(itemsTyped.length);
const sortBy = ref<BTableSortBy[]>([]);
const sortDirection = ref("asc");
const filter = ref("");
const filterOn = ref([]);

function onFiltered(filteredItems: TableItem<Order>[]) {
  // Trigger pagination to update the number of buttons/pages due to filtering
  totalRows.value = filteredItems.length;
  currentPage.value = 1;
}
</script>
<style>
.list-group {
  max-height: 80vh;
}
.cardin {
  overflow-y: auto;
  overflow-x: hidden;
  scrollbar-width: thin;
  scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
  scroll-margin: 0;
  scroll-padding: 0;
  scrollbar-gutter: 0;
}
.order-table {
  overflow-y: auto;
  overflow-x: hidden;

  scrollbar-width: thin;
  scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
}
</style>
