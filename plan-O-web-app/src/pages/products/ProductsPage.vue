<template>
  <div
    class="col-auto col-xxl-10 col-xl-10 col-md-9 col-sm-9 align-items-start mx-0 px-0 mh-100"
  >
    <BNavbar
      toggleable="expand-lg"
      variant="dark bg-opacity-75"
      v-b-color-mode="'dark'"
      class="align-self-start align-items-center shadow-lg"
    >
      <BNavbarNav>
        <BBreadcrumb
          class="text-decoration-none"
          :items="breadcrumbStringArray"
        />
      </BNavbarNav>
      <BNavbarNav class="w-50">
        <BNavForm class="d-flex w-100">
          <BFormInput id="input-1" class="mx-2" placeholder="Поиск..." />
          <BButton type="submit" variant="outline-light">Искать</BButton>
        </BNavForm>
      </BNavbarNav>
    </BNavbar>

    <BContainer class="py-3 px-3 mx-0" :fluid="true">
      <div class="d-flex mb-4">
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
        <!-- FILTER -->
        <!-- <BFormGroup
                        v-slot="{ariaDescribedby}"
                        v-model="sortDirection"
                        label="Filter On"
                        description="Leave all unchecked to filter on all data"
                        label-cols-sm="3"
                        label-align-sm="right"
                        label-size="sm"
                        class="mb-0"
                        >
                        <div class="d-flex gap-2">
                            <BFormCheckbox v-model="filterOn" value="name" :aria-describedby="ariaDescribedby"
                            >Name</BFormCheckbox
                            >
                            <BFormCheckbox v-model="filterOn" value="age" :aria-describedby="ariaDescribedby"
                            >Age</BFormCheckbox
                            >
                            <BFormCheckbox v-model="filterOn" value="isActive" :aria-describedby="ariaDescribedby"
                            >Active</BFormCheckbox
                            >
                        </div>
                        </BFormGroup> -->
      </div>

      <!-- Product TABLE-->
      <BTable
        class="Product-table mt-1 rounded-3"
        style="max-height: 75vh; max-width: 100%"
        model="sortBy"
        :sort-internal="true"
        :items="itemsTyped"
        :fields="fieldsTyped"
        :current-page="currentPage"
        :per-page="perPage"
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
            variant="light text-dark"
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
              /></svg>Просмотр</BButton
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
              viewBox="0 0 16 16"
            >
              <path
                d="M12.854.146a.5.5 0 0 0-.707 0L10.5 1.793 14.207 5.5l1.647-1.646a.5.5 0 0 0 0-.708zm.646 6.061L9.793 2.5 3.293 9H3.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.207zm-7.468 7.468A.5.5 0 0 1 6 13.5V13h-.5a.5.5 0 0 1-.5-.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.5-.5V10h-.5a.5.5 0 0 1-.175-.032l-.179.178a.5.5 0 0 0-.11.168l-2 5a.5.5 0 0 0 .65.65l5-2a.5.5 0 0 0 .168-.11z"
              />
            </svg>
            Изменить
          </BButton>
        </template>
      </BTable>
      <BRow class="my-2 align-self-bottom">
        <!-- PAGINATION -->
        <BCol sm="6">
          <BPagination
            v-model="currentPage"
            :total-rows="totalRows"
            :per-page="perPage"
            :align="'end'"
            size="sm"
            class="my-0"
          />
        </BCol>
        <BCol sm="6">
          <BFormGroup
            label="Per page"
            label-for="per-page-select"
            label-cols-sm="10"
            label-align-sm="end"
            label-size="sm"
            class="mb-0 d-flex justify-content-end"
          >
            <BFormSelect
              id="per-page-select"
              v-model="perPage"
              :options="pageOptions"
              size="sm"
            />
          </BFormGroup>
        </BCol>
      </BRow>

      <!-- Info modal -->
      <BModal
        size="lg"
        :id="infoProduct.id"
        v-model="infoProduct.open"
        :title="infoProduct.title"
        :ok-only="true"
        @hide="resetInfoProduct"
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
                v-model="infoProduct.content.productName"
                type="email"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextProduct" variant="outline-dark"
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
                v-model="infoProduct.content.email"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextProduct(info)" variant="outline-dark"
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
                v-model="infoProduct.content.productType"
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
                v-model="infoProduct.content.adress"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextProduct" variant="outline-dark"
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
                v-model="infoProduct.content.phone"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextProduct" variant="outline-dark"
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
                v-model="infoProduct.content.managerName"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextProduct" variant="outline-dark"
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
                v-model="infoProduct.content.comments"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextProduct" variant="outline-dark"
                  >Copy</BButton
                >
              </BInputGroupAppend>
            </BInputGroup>
          </BFormGroup>
        </BForm>
      </BModal>

      <!-- Edit modal -->
      <BModal
        size="lg"
        :id="editProduct.id"
        v-model="editProduct.open"
        :title="editProduct.title"
        hideFooter="true"
        @hide="resetEditProduct"
      >
        <BForm @submit="onSubmit(edit.item)" v-if="show">
          <BFormGroup
            id="edit-input-group-1"
            label="Название компании"
            label-for="edit-input-1"
            class="mb-1"
          >
            <BFormInput
              id="edit-input-1"
              v-model="editProduct.content.productName"
              type="text"
              required
            />
          </BFormGroup>
          <BFormGroup
            id="edit-input-group-2"
            label="Email"
            label-for="edit-input-2"
            class="my-1"
          >
            <BFormInput
              id="edit-input-2"
              v-model="editProduct.content.email"
              type="email"
            />
          </BFormGroup>
          <BFormGroup
            id="input-group-3"
            label="Тип клиента"
            label-for="edit-input-3"
            class="my-1"
          >
            <BFormSelect
              id="edit-input-3"
              v-model="editProduct.content.productType"
              :options="productTypes"
              required
            />
          </BFormGroup>
          <BFormGroup
            id="edit-input-group-4"
            label="Адрес"
            label-for="edit-input-4"
            class="my-1"
          >
            <BFormInput
              id="edit-input-4"
              v-model="editProduct.content.adress"
              type="text"
            />
          </BFormGroup>
          <BFormGroup
            id="edit-input-group-5"
            label="Номер телефона"
            label-for="edit-input-5"
            description="например, +7 (911) 007-31-94"
            class="my-1"
          >
            <BFormInput
              id="edit-input-5"
              v-model="editProduct.content.phone"
              type="tel"
              masked="true"
              mask="+!!! (###) ###-##-##"
            />
          </BFormGroup>
          <BFormGroup
            id="edit-input-group-6"
            label="Имя менеджера"
            label-for="edit-input-6"
            class="my-1"
          >
            <BFormInput
              id="edit-input-6"
              v-model="editProduct.content.managerName"
              type="text"
            />
          </BFormGroup>
          <BFormGroup
            id="edit-input-group-7"
            label="Комментарии"
            label-for="edit-input-7"
            class="my-1"
          >
            <BFormInput
              id="edit-input-7"
              v-model="editProduct.content.comments"
              type="text"
            />
          </BFormGroup>
          <BButton type="submit" variant="success" class="mt-3 me-2"
            >Сохранить</BButton
          >
          <BButton
            type="submit"
            variant="outline-danger"
            class="mt-3"
            @click="
              () => {
                editProduct.open = 'false';
              }
            "
            >Отмена</BButton
          >
        </BForm>
      </BModal>

      <!-- Add modal -->
      <BModal
        size="lg"
        :id="addProduct.id"
        v-model="addProduct.open"
        :title="addProduct.title"
        hideFooter="true"
        @hide="resetAddProduct"
      >
        <BForm @submit="onSubmit()" v-if="show">
          <BFormGroup
            id="add-input-group-1"
            label="Название товара"
            label-for="add-input-1"
            class="mb-1"
          >
            <BFormInput
              id="add-input-1"
              v-model="addProduct.content.productName"
              type="text"
              required
            />
          </BFormGroup>
          <BFormGroup
            id="input-group-2"
            label="Тип товара"
            label-for="add-input-2"
            class="my-1"
          >
            <BFormSelect
              id="add-input-2"
              v-model="addProduct.content.productType"
              :options="productTypes"
              required
            />
          </BFormGroup>
          <BFormGroup
            id="add-input-group-3"
            label="Рекомендованная цена"
            label-for="add-input-3"
            class="my-1"
          >
            <BFormInput
              id="add-input-3"
              v-model="addProduct.content.productPrice"
              type="number"
            />
          </BFormGroup>
          <BFormGroup
            id="add-input-group-4"
            label="Единица измерения"
            label-for="add-input-4"
            class="my-1"
          >
            <BFormSelect
              id="add-input-4"
              v-model="addProduct.content.productUnit"
              :options="productUnits"
              required
            />
          </BFormGroup>
          <BButton type="submit" variant="success" class="mt-3 me-2"
            >Сохранить</BButton
          >
          <BButton
            type="submit"
            variant="outline-danger"
            class="mt-3"
            @click="
              () => {
                addProduct.open = 'false';
              }
            "
            >Отмена</BButton
          >
        </BForm>
      </BModal>
    </BContainer>
  </div>
</template>

<script setup lang="ts">

import {
  BButton,
  BFormSelect,
  BInputGroup,
  BFormCheckbox,
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

interface Product {
  productName: string;
  productPrice: string;
  productType: string;
  productUnit: string;
}
const itemsTyped: TableItem<Product>[] = [];

const editProduct = reactive({
  open: false,
  id: "edit-modal",
  title: "",
  content: {
    productName: "",
    productPrice: "",
    productType: "",
    productUnit: "",
  },
});

const addProduct = reactive({
  open: false,
  id: "add-modal",
  title: "",
  content: {
    productName: "",
    productPrice: "",
    productType: "",
    productUnit: "",
  },
});

const show = ref(true);

const onSubmit = (item) => {
  alert(JSON.stringify(item));
};
const breadcrumbStringArray = ["CRM", "Товары"];

const fieldsTyped: Exclude<TableFieldRaw<Product>, string>[] = [
  {
    key: "productName",
    label: "Название",
    sortable: true,
    sortDirection: "desc",
    class: "text-sm",
  },
  {
    key: "productType",
    label: "Тип товара",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "productPrice",
    label: "Рекомендованная цена",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "productUnit",
    label: "Единица измерения",
    sortable: true,
    sortDirection: "desc",
  },
  { key: "actions", label: "Действия" },
];

const pageOptions = [
  { value: 5, text: "5" },
  { value: 10, text: "10" },
  { value: 15, text: "15" },
  { value: 30, text: "30" },
  { value: 100, text: "100" },
];

const totalRows = ref(itemsTyped.length);
const currentPage = ref(1);
const perPage = ref(30);
const sortBy = ref<BTableSortBy[]>([]);
const sortDirection = ref("asc");
const filter = ref("");
const filterOn = ref([]);

const productTypes = [
  { text: "Выберите..", value: null },
  "Наклейки",
  "Объемные наклейки",
  "ABS-пластик",
  "Металлизированные изделия",
  "Пластик",
  "Шелкография",
];

const productUnits = [{ text: "Выберите..", value: null }, "Кг", "Л"];
const infoProduct = reactive({
  open: false,
  id: "info-modal",
  title: "",
  content: {},
});

// Create an options list from our fields
function info(item: TableItem<Product>, index: number) {
  infoProduct.title = `${item.productName}`;
  infoProduct.content = item;
  infoProduct.open = true;
}

function edit(item: TableItem<Product>, index: number) {
  editProduct.title = `${item.productName}`;
  editProduct.open = true;
  editProduct.content.productName = item.productName;
  editProduct.content.productType = item.productType;
  editProduct.content.productPrice = item.productPrice;
  editProduct.content.productUnit = item.productUnit;
}
function add() {
  addProduct.title = `Новый клиент`;
  addProduct.open = true;
  addProduct.content.productName = "";
  addProduct.content.productType = "";
  addProduct.content.productPrice = "";
  addProduct.content.productUnit = "";
}

function resetInfoProduct() {
  infoProduct.title = "";
  infoProduct.content = "";
}
function resetEditProduct() {
  infoProduct.title = "";
  infoProduct.content = "";
}
function resetAddProduct() {
  infoProduct.title = "";
  infoProduct.content = "";
}

function copyTextProduct() {}

function onFiltered(filteredItems: TableItem<Product>[]) {
  // Trigger pagination to update the number of buttons/pages due to filtering
  totalRows.value = filteredItems.length;
  currentPage.value = 1;
}
</script>

<style>
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
.Product-table {
  overflow-y: auto;
  overflow-x: hidden;

  scrollbar-width: thin;
  scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
}
</style>
