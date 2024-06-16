<template>
  <BContainer class="p-2 mx-0" :fluid="true">
    <div class="d-flex mb-4">
      <BButton type="submit" class="text-info" variant="dark me-2" @click="add">
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

    <!-- Supplier TABLE-->
    <BTable
      class="Supplier-table mt-1 rounded-3"
      style="max-height: calc(100vh - 13em); max-width: 100%"
      model="sortBy"
        :sort-internal="true"
        :items="itemsTyped"
        :fields="fieldsTyped"
        :filter="filter"
        :responsive="false"
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

    <!-- Info modal -->
    <BModal
      class="info"
      size="lg"
      :id="infoSupplier.id"
      v-model="infoSupplier.open"
      :title="infoSupplier.title"
      :ok-only="true"
      @hide="resetInfoSupplier"
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
              v-model="infoSupplier.content.companyName"
              type="text"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.companyName)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="info-input-group-2"
          label="Поставляемые товары"
          label-for="info-input-2"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-2"
              v-model="infoSupplier.content.suppliableProductType"
              type="text"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.suppliableProductType)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="info-input-group-3"
          label="Имя менеджера"
          label-for="info-input-3"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-3"
              v-model="infoSupplier.content.managerName"
              type="text"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.managerName)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="info-input-group-4"
          label="Название компании"
          label-for="info-input-4"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-4"
              v-model="infoSupplier.content.email"
              type="email"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.email)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="info-input-group-5"
          label="Телефон"
          label-for="info-input-5"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-5"
              v-model="infoSupplier.content.phone"
              type="text"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.phone)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="info-input-group-6"
          label="Веб-сайт"
          label-for="info-input-6"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-6"
              v-model="infoSupplier.content.websiteLink"
              type="url"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.websiteLink)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="info-input-group-7"
          label="Тип оплаты"
          label-for="info-input-7"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-7"
              v-model="infoSupplier.content.paymentType"
              type="url"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.paymentType)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="info-input-group-8"
          label="Комментарии"
          label-for="info-input-8"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-8"
              v-model="infoSupplier.content.comments"
              type="text"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyText(infoSupplier.content.comments)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>
      </BForm>
    </BModal>

    <!-- Edit modal -->
    <BModal
      class="edit"
      size="lg"
      :id="editSupplier.id"
      v-model="editSupplier.open"
      :title="editSupplier.title"
      :ok-only="true"
      @hide="resetEditSupplier"
      hideFooter= true
    >
      <BForm 
      v-if="show"
      @submit="onSubmitEdit()">
        <!-- Название компании -->
        <BFormGroup
          id="edit-input-group-1"
          label="Название компании"
          label-for="edit-input-1"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-1"
              v-model="editSupplier.content.companyName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="edit-input-group-2"
          label="Поставляемые товары"
          label-for="edit-input-2"
          class="mb-1"
        >
          <BInputGroup>
            <BFormSelect
            id="edit-input-2"
            v-model="editSupplier.content.suppliableProductType"
            :options="suppliableSupplierTypes"
            required
          />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="edit-input-group-3"
          label="Имя менеджера"
          label-for="edit-input-3"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-3"
              v-model="editSupplier.content.managerName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="edit-input-group-4"
          label="E-mail"
          label-for="edit-input-4"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-4"
              v-model="editSupplier.content.email"
              type="email"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="edit-input-group-5"
          label="Телефон"
          label-for="edit-input-5"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-5"
              v-model="editSupplier.content.phone"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="edit-input-group-6"
          label="Веб-сайт"
          label-for="edit-input-6"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-6"
              v-model="editSupplier.content.websiteLink"
              type="url"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="edit-input-group-7"
          label="Тип оплаты"
          label-for="edit-input-7"
          class="mb-1"
        >
          <BInputGroup>
            <BFormSelect
            id="edit-input-7"
            v-model="editSupplier.content.paymentType"
            :options="paymentTypes"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="edit-input-group-8"
          label="Комментарии"
          label-for="edit-input-8"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-8"
              v-model="editSupplier.content.comments"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton type="submit" variant="success" class="mt-3 me-2"
          >Сохранить</BButton
        >
        <BButton
          type="submit"
          variant="outline-danger"
          class="mt-3"
          @click="resetEditClient"
          >Отмена</BButton
        >
      </BForm>
    </BModal>

    <!-- Add modal -->
    <BModal
      class="add"
      size="lg"
      :id="addSupplier.id"
      v-model="addSupplier.open"
      :title="addSupplier.title"
      :ok-only="true"
      @hide="resetAddSupplier"
      hideFooter=true
    >
      <BForm 
      v-if="show"
      @submit="onSubmitAdd(addSupplier.content)">
        <!-- Название компании -->
        <BFormGroup
          id="add-input-group-1"
          label="Название компании"
          label-for="add-input-1"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-1"
              v-model="addSupplier.content.companyName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="add-input-group-2"
          label="Поставляемые товары"
          label-for="add-input-2"
          class="mb-1"
        >
        <BInputGroup>
            <BFormSelect
            id="add-input-2"
            v-model="addSupplier.content.suppliableProductType"
            :options="suppliableSupplierTypes"
            required
          />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="add-input-group-3"
          label="Имя менеджера"
          label-for="add-input-3"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-3"
              v-model="addSupplier.content.managerName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="add-input-group-4"
          label="E-mail"
          label-for="add-input-4"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-4"
              v-model="addSupplier.content.email"
              type="email"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="add-input-group-5"
          label="Телефон"
          label-for="add-input-5"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-5"
              v-model="addSupplier.content.phone"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="add-input-group-6"
          label="Веб-сайт"
          label-for="add-input-6"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-6"
              v-model="addSupplier.content.websiteLink"
              type="url"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="add-input-group-7"
          label="Тип оплаты"
          label-for="add-input-7"
          class="mb-1"
        >
        <BInputGroup>
            <BFormSelect
            id="edit-input-7"
            v-model="addSupplier.content.paymentType"
            :options="paymentTypes"
            />
          </BInputGroup>
        </BFormGroup>
        <BFormGroup
          id="add-input-group-8"
          label="Комментарии"
          label-for="add-input-8"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-8"
              v-model="addSupplier.content.comments"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton type="submit" variant="success" class="mt-3 me-2"
          >Сохранить</BButton
        >
        <BButton
          type="submit"
          variant="outline-danger"
          class="mt-3"
          @click="resetaddClient()"
          >Отмена</BButton
        >
      </BForm>
    </BModal>
  </BContainer>
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

interface Supplier {
  companyName: string;
  managerName: string;
  email: string;
  phone: string;
  paymentType: string;
  suppliableProductType: string;
  websiteLink: string;
  comments: string;
}

const suppliableSupplierTypes = [
  { text: "Выберите..", value: null },
  "смола",
  "пленка",
  "запчасти",
  "промышленная химия",
  "слесарно-токарные запчасти",
  "клапаны",
  "шелкография",
  "посредники",
];
const paymentTypes = ["Предоплата", "50/50", "Постоплата"];

const itemsTyped: TableItem<Supplier>[] = reactive([
  {
    companyName: "Resin Solutions",
    managerName: "Ivan Ivanov",
    email: "ivan.ivanov@resinsolutions.com",
    phone: "+7 495 123 4567",
    paymentType: "Предоплата",
    suppliableProductType: "смола",
    websiteLink: "http://resinsolutions.com",
    comments: "Leading supplier in resin Suppliers.",
  },
  {
    companyName: "FilmFlex",
    managerName: "Olga Petrov",
    email: "olga.petrov@filmflex.com",
    phone: "+7 812 987 6543",
    paymentType: "50/50",
    suppliableProductType: "пленка",
    websiteLink: "http://filmflex.com",
    comments: "High-quality film supplier.",
  },
  {
    companyName: "Spare Parts Hub",
    managerName: "Alexander Sidorov",
    email: "alexander.sidorov@sparepartshub.com",
    phone: "+7 495 765 4321",
    paymentType: "Постоплата",
    suppliableProductType: "запчасти",
    websiteLink: "http://sparepartshub.com",
    comments: "Wide range of spare parts available.",
  },
  {
    companyName: "Industrial Chem",
    managerName: "Nikolay Smirnov",
    email: "nikolay.smirnov@industrialchem.com",
    phone: "+7 499 123 7890",
    paymentType: "Предоплата",
    suppliableProductType: "промышленная химия",
    websiteLink: "http://industrialchem.com",
    comments: "Specializing in industrial chemicals.",
  },
  {
    companyName: "Machinist Supplies",
    managerName: "Maria Ivanova",
    email: "maria.ivanova@machinistsupplies.com",
    phone: "+7 812 345 6789",
    paymentType: "50/50",
    suppliableProductType: "слесарно-токарные запчасти",
    websiteLink: "http://machinistsupplies.com",
    comments: "Precision machinist parts supplier.",
  },
  {
    companyName: "Valve Experts",
    managerName: "Dmitry Petrov",
    email: "dmitry.petrov@valveexperts.com",
    phone: "+7 495 234 5678",
    paymentType: "Постоплата",
    suppliableProductType: "клапаны",
    websiteLink: "http://valveexperts.com",
    comments: "Top quality valves for all needs.",
  },
  {
    companyName: "Silk Screen Prints",
    managerName: "Tatiana Sidorova",
    email: "tatiana.sidorova@silkscreenprints.com",
    phone: "+7 812 876 5432",
    paymentType: "Предоплата",
    suppliableProductType: "шелкография",
    websiteLink: "http://silkscreenprints.com",
    comments: "Expert in silk screen printing.",
  },
  {
    companyName: "Mediation Services",
    managerName: "Sergey Smirnov",
    email: "sergey.smirnov@mediationservices.com",
    phone: "+7 499 321 0987",
    paymentType: "50/50",
    suppliableProductType: "посредники",
    websiteLink: "http://mediationservices.com",
    comments: "Trusted intermediary services.",
  },
  {
    companyName: "Polymer World",
    managerName: "Elena Ivanova",
    email: "elena.ivanova@polymerworld.com",
    phone: "+7 495 543 2109",
    paymentType: "Постоплата",
    suppliableProductType: "смола",
    websiteLink: "http://polymerworld.com",
    comments: "Global supplier of polymers.",
  },
  {
    companyName: "FilmCo",
    managerName: "Alexey Petrov",
    email: "alexey.petrov@filmco.com",
    phone: "+7 812 678 1234",
    paymentType: "Предоплата",
    suppliableProductType: "пленка",
    websiteLink: "http://filmco.com",
    comments: "High-performance film Suppliers.",
  },
  {
    companyName: "SpareParts Ltd.",
    managerName: "Irina Sidorova",
    email: "irina.sidorova@sparepartsltd.com",
    phone: "+7 495 789 4321",
    paymentType: "50/50",
    suppliableProductType: "запчасти",
    websiteLink: "http://sparepartsltd.com",
    comments: "Reliable spare parts supplier.",
  },
  {
    companyName: "ChemPro",
    managerName: "Vladimir Smirnov",
    email: "vladimir.smirnov@chempro.com",
    phone: "+7 499 210 9876",
    paymentType: "Постоплата",
    suppliableProductType: "промышленная химия",
    websiteLink: "http://chempro.com",
    comments: "Chemical Suppliers for industries.",
  },
]);

const show = ref(true);

const breadcrumbStringArray = ["CRM", "Товары"];

const fieldsTyped: Exclude<TableFieldRaw<Supplier>, string>[] = [
  {
    key: "companyName",
    label: "Поставщик",
    sortable: true,
    sortDirection: "desc",
    class: "text-sm",
  },
  {
    key: "suppliableProductType",
    label: "Поставляемые товары",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "managerName",
    label: "Имя менеджера",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "email",
    label: "E-mail",
    sortable: true,
    sortDirection: "desc",
  },
  {
    key: "phone",
    label: "Телефон",
    sortable: false,
  },
  {
    key: "websiteLink",
    label: "Веб-сайт",
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
const filter = ref("");
const filterOn = ref([]);

// INFO
function info(item: TableItem<Supplier>, index: number) {
  infoSupplier.title = `${item.companyName}`;
  infoSupplier.content = { ...item };
  infoSupplier.open = true;
}

const infoSupplier = reactive({
  open: false,
  id: "info-modal",
  title: "",
  content: {},
});

function resetInfoSupplier() {
  infoSupplier.title = "";
  infoSupplier.content = "";
  infoSupplier.open = false;
}

// EDIT
function edit(item: TableItem<Supplier>, index: number) {
  editSupplier.title = `${item.companyName}`;
  editSupplier.open = true;
  editSupplier.content = item;
}

const editSupplier = reactive({
  open: false,
  id: "edit-modal",
  title: "",
  content: {},
});

const resetEditSupplier = () => {
  editSupplier.title = "";
  editSupplier.content = {};
  editSupplier.open = false;
};
function onSubmitEdit() {
  resetEditSupplier();
}

// ADD
function add() {
  addSupplier.title = `Новый поставщик`;
  addSupplier.open = true;
  addSupplier.content = {
    companyName: "",
    managerName: "",
    email: "",
    phone: "",
    paymentType: "",
    suppliableSupplier: "",
    websiteLink: "",
    comments: "",
  };
}

const addSupplier = reactive({
  open: false,
  id: "add-modal",
  title: "",
  content: {
    companyName: "",
    managerName: "",
    email: "",
    phone: "",
    paymentType: "",
    suppliableSupplier: "",
    websiteLink: "",
    comments: "",
  },
});

function resetAddSupplier() {
  addSupplier.title = "";
  addSupplier.content = {
    companyName: "",
    managerName: "",
    email: "",
    phone: "",
    paymentType: "",
    suppliableSupplier: "",
    websiteLink: "",
    comments: ""
  };
  addSupplier.open = false;
}

const onSubmitAdd = (item) => {
  itemsTyped.push(item);
  resetAddSupplier();
};
function copyText(input) {
  navigator.clipboard.writeText(input)
}

function onFiltered(filteredItems: TableItem<Supplier>[]) {
  // Trigger pagination to update the number of buttons/pages due to filtering
  totalRows.value = filteredItems.length;
  currentPage.value = 1;
}
</script>

<style>
.Supplier-table {
  overflow-y: auto;
  overflow-x: hidden;

  scrollbar-width: thin;
  scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
}
</style>
