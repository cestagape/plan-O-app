<template>
  
    <BContainer class="py-2 px-2 mx-0" :fluid="true">
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
      </div>

      <!-- CLIENT TABLE-->
      <BTable
        class="client-table mt-1 rounded-3"
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
        ref="table"
        id="my-table"
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
              /></svg>Изменить</BButton
          >
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
        :id="infoClient.id"
        v-model="infoClient.open"
        :title="infoClient.title"
        :ok-only="true"
        @hide="resetInfoClient"
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
                v-model="infoClient.content.companyName"
                type="email"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextClient" variant="outline-dark"
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
                v-model="infoClient.content.email"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextClient(info)" variant="outline-dark"
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
                v-model="infoClient.content.customerType"
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
                v-model="infoClient.content.adress"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextClient" variant="outline-dark"
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
                v-model="infoClient.content.phone"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextClient" variant="outline-dark"
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
                v-model="infoClient.content.managerName"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextClient" variant="outline-dark"
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
                v-model="infoClient.content.comments"
                type="text"
                disabled
              />
              <BInputGroupAppend>
                <BButton @click="copyTextClient" variant="outline-dark"
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
        :id="editClient.id"
        v-model="editClient.open"
        :title="editClient.title"
        hideFooter="true"
        @hide="resetEditClient"
      >
        <BForm @submit="onSubmitEdit()" v-if="show">
          <BFormGroup
            id="edit-input-group-1"
            label="Название компании"
            label-for="edit-input-1"
            class="mb-1"
          >
            <BFormInput
              id="edit-input-1"
              v-model="editClient.content.companyName"
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
              v-model="editClient.content.email"
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
              v-model="editClient.content.customerType"
              :options="clientTypes"
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
              v-model="editClient.content.adress"
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
              v-model="editClient.content.phone"
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
              v-model="editClient.content.managerName"
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
              v-model="editClient.content.comments"
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
                editClient.open = 'false';
              }
            "
            >Отмена</BButton
          >
        </BForm>
      </BModal>

      <!-- Add modal -->
      <BModal
        size="lg"
        :id="addClient.id"
        v-model="addClient.open"
        :title="addClient.title"
        hideFooter="true"
        @hide="resetAddClient"
      >
        <BForm  
        v-if="show"
        @submit="onSubmitAdd(addClient.content)">
          <BFormGroup
            id="add-input-group-1"
            label="Название компании"
            label-for="add-input-1"
            class="mb-1"
          >
            <BFormInput
              id="add-input-1"
              v-model="addClient.content.companyName"
              type="text"
              required
            />
          </BFormGroup>
          <BFormGroup
            id="add-input-group-2"
            label="Email"
            label-for="add-input-2"
            class="my-1"
          >
            <BFormInput
              id="add-input-2"
              v-model="addClient.content.email"
              type="email"
            />
          </BFormGroup>
          <BFormGroup
            id="input-group-3"
            label="Тип клиента"
            label-for="add-input-3"
            class="my-1"
          >
            <BFormSelect
              id="add-input-3"
              v-model="addClient.content.customerType"
              :options="clientTypes"

            />
          </BFormGroup>
          <BFormGroup
            id="add-input-group-4"
            label="Адрес"
            label-for="add-input-4"
            class="my-1"
          >
            <BFormInput
              id="add-input-4"
              v-model="addClient.content.adress"
              type="text"
            />
          </BFormGroup>
          <BFormGroup
            id="add-input-group-5"
            label="Номер телефона"
            label-for="add-input-5"
            description="например, +7 (911) 007-31-94"
            class="my-1"
          >
            <BFormInput
              id="add-input-5"
              v-model="addClient.content.phone"
              type="tel"
              masked="true"
              mask="+!!! (###) ###-##-##"
            />
          </BFormGroup>
          <BFormGroup
            id="add-input-group-6"
            label="Имя менеджера"
            label-for="add-input-6"
            class="my-1"
          >
            <BFormInput
              id="add-input-6"
              v-model="addClient.content.managerName"
              type="text"
            />
          </BFormGroup>
          <BFormGroup
            id="add-input-group-7"
            label="Имя менеджера"
            label-for="add-input-7"
            class="my-1"
          >
            <BFormInput
              id="add-input-7"
              v-model="addClient.content.comments"
              type="text"
            />
          </BFormGroup>
          <BButton 
          type="submit"
          variant="success" 
          class="mt-3 me-2"
          >
            Сохранить</BButton>
          <BButton
            variant="outline-danger"
            class="mt-3"
            @click="resetAddClient()"
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
  type TableFieldRaw,
  type TableItem,
} from "bootstrap-vue-next";
import { reactive, ref, getCurrentInstance } from "vue";



interface Client {
  companyName: string;
  customerType: string;
  adress: string;
  email: string;
  phone: string;
  managerName: string;
  comments: string;
};

const itemsTyped = reactive([
  {
    "companyName": "BlueSky Solutions",
    "customerType": "Corporate",
    "adress": "123 Oak St, Springfield, IL",
    "email": "contact@bluesky.com",
    "phone": "555-123-4567",
    "managerName": "John Smith",
    "comments": "Looking for long-term partnership"
  },
  {
    "companyName": "GreenLeaf Industries",
    "customerType": "Manufacturing",
    "adress": "456 Maple Ave, Springfield, IL",
    "email": "info@greenleaf.com",
    "phone": "555-234-5678",
    "managerName": "Emily Johnson",
    "comments": "Interested in expanding services"
  },
  {
    "companyName": "TechWave Innovations",
    "customerType": "Technology",
    "adress": "789 Cedar Blvd, Springfield, IL",
    "email": "sales@techwave.com",
    "phone": "555-345-6789",
    "managerName": "Michael Brown",
    "comments": "Project timeline is flexible"
  },
  {
    "companyName": "RedRiver Logistics",
    "customerType": "Logistics",
    "adress": "101 River St, Springfield, IL",
    "email": "support@redriver.com",
    "phone": "555-456-7890",
    "managerName": "Sarah Wilson",
    "comments": "Looking for bulk shipping rates"
  },
  {
    "companyName": "PrimeHealth Solutions",
    "customerType": "Healthcare",
    "adress": "234 Elm St, Springfield, IL",
    "email": "office@primehealth.com",
    "phone": "555-567-8901",
    "managerName": "David Martinez",
    "comments": "Needs flexible scheduling options"
  },
  {
    "companyName": "Pinnacle Realty",
    "customerType": "Real Estate",
    "adress": "567 Birch Ave, Springfield, IL",
    "email": "inquiries@pinnaclerealty.com",
    "phone": "555-678-9012",
    "managerName": "Olivia Adams",
    "comments": "Looking for property management"
  },
  {
    "companyName": "Skyline Constructions",
    "customerType": "Construction",
    "adress": "890 Pine Blvd, Springfield, IL",
    "email": "projects@skyline.com",
    "phone": "555-789-0123",
    "managerName": "Christopher Lee",
    "comments": "Needs estimates for new project"
  },
  {
    "companyName": "Horizon Financial",
    "customerType": "Finance",
    "adress": "123 Elmwood Ln, Springfield, IL",
    "email": "accounts@horizon.com",
    "phone": "555-890-1234",
    "managerName": "Sophia Ramirez",
    "comments": "Wants to discuss investment options"
  },
  {
    "companyName": "Innovatech Labs",
    "customerType": "Research",
    "adress": "456 Sycamore Rd, Springfield, IL",
    "email": "research@innovatech.com",
    "phone": "555-901-2345",
    "managerName": "Daniel Carter",
    "comments": "Exploring new collaborative projects"
  },
  {
    "companyName": "PureNature Organics",
    "customerType": "Agriculture",
    "adress": "789 Willow St, Springfield, IL",
    "email": "sales@purenature.com",
    "phone": "555-012-3456",
    "managerName": "Chloe Green",
    "comments": "Interested in bulk organic produce"
  },
  {
    "companyName": "Peak Fitness",
    "customerType": "Health and Wellness",
    "adress": "101 Park Ave, Springfield, IL",
    "email": "membership@peakfitness.com",
    "phone": "555-123-9876",
    "managerName": "Jacob Harris",
    "comments": "Wants to arrange fitness seminars"
  },
  {
    "companyName": "Sunrise Travel",
    "customerType": "Travel",
    "adress": "234 Beach St, Springfield, IL",
    "email": "info@sunrisetravel.com",
    "phone": "555-234-9876",
    "managerName": "Isabella Wright",
    "comments": "Looking for group travel packages"
  },
  {
    "companyName": "NorthStar Education",
    "customerType": "Education",
    "adress": "567 College Ave, Springfield, IL",
    "email": "admissions@northstar.edu",
    "phone": "555-345-9876",
    "managerName": "Liam Patterson",
    "comments": "Discussing internship opportunities"
  },
  {
    "companyName": "Dynamic Design Studio",
    "customerType": "Design",
    "adress": "890 Art St, Springfield, IL",
    "email": "designs@dynamic.com",
    "phone": "555-456-9876",
    "managerName": "Emma Walker",
    "comments": "Exploring branding options"
  },
  {
    "companyName": "Evergreen Landscaping",
    "customerType": "Landscaping",
    "adress": "101 Green Way, Springfield, IL",
    "email": "contact@evergreen.com",
    "phone": "555-567-9876",
    "managerName": "Alexander Rivera",
    "comments": "Interested in seasonal services"
  },
  {
    "companyName": "Visionary Marketing",
    "customerType": "Marketing",
    "adress": "234 Market St, Springfield, IL",
    "email": "inquiries@visionary.com",
    "phone": "555-678-9876",
    "managerName": "Ava Bennett",
    "comments": "Exploring new digital strategies"
  },
  {
    "companyName": "GlobalTech Systems",
    "customerType": "Information Technology",
    "adress": "567 Silicon Blvd, Springfield, IL",
    "email": "support@globaltech.com",
    "phone": "555-789-9876",
    "managerName": "William Sanchez",
    "comments": "Needs assistance with cloud services"
  },
  {
    "companyName": "TopShelf Catering",
    "customerType": "Catering",
    "adress": "890 Gourmet Ln, Springfield, IL",
    "email": "catering@topshelf.com",
    "phone": "555-890-9876",
    "managerName": "Sophia Turner",
    "comments": "Looking for event catering options"
  },
])
const show = ref(true);

const breadcrumbStringArray = ["CRM", "Клиенты"];



const fieldsTyped: Exclude<TableFieldRaw<Client>, string>[] = [
  {
    key: "companyName",
    label: "Название",
    sortable: true,
    sortDirection: "desc",
    class: "text-sm",
  },
  {
    key: "customerType",
    label: "Тип клиента",
    sortable: true,
    sortDirection: "desc",
  },
  { key: "email", label: "Email", sortable: false, sortDirection: "desc" },
  {
    key: "phone",
    label: "Номер телефона",
    sortable: false,
    sortDirection: "desc",
    class: "",
  },
  {
    key: "managerName",
    label: "Имя менеджера",
    sortable: false,
    sortDirection: "desc",
  },
  {
    key: "comments",
    label: "Комментарии",
    sortable: false,
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
const filter = ref(null);
const filterOn = ref([]);

const clientTypes = [
  { text: "Выберите..", value: null },
  "Частный",
  "Корпорация",
  "Самозанятый",
];

const infoClient = reactive({
  open: false,
  id: "info-modal",
  title: "",
  content: {},
});

const addClient = reactive({
  open: false,
  id: "add-modal",
  title: "",
  content: {},
});

const editClient = reactive({
  open: false,
  id: "edit-modal",
  title: "",
  content: {},
  item: {},
});

// Create an options list from our fields
function info(item: TableItem<Client>, index: number) {
  infoClient.title = `${item.companyName}`;
  infoClient.content = { ...item }
  infoClient.open = true;
}

function edit(item: TableItem<Client>, index: number) {
  editClient.title = `${item.companyName}`;
  editClient.open = true;
  editClient.content = item
  editClient.item = item

}
function add() {
  addClient.title = `Новый клиент`;
  addClient.open = true;
  addClient.content = {
    companyName: "",
    customerType: "",
    adress: "",
    email: "",
    phone: "",
    managerName: "",
    comments: "",
  }
}


const onSubmitAdd = (obj) => {
  itemsTyped.push(obj);
  addClient.open = false;
  // TODO:
  //  Добавить метод по отправке даннных на сервер в 

};
const onSubmitEdit = () => {
  // alert(JSON.stringify(item))
  editClient.item = { ...editClient.content }
  editClient.open = false;
  // TODO:
  //  Добавить метод по отправке даннных на сервер в 
};

function resetInfoClient() {
  infoClient.title = "";
  infoClient.content = "";
  infoClient.open = false;
}
function resetEditClient() {
  editClient.title = "";
  editClient.content = {
    companyName: "",
    customerType: "",
    adress: "",
    email: "",
    phone: "",
    managerName: "",
    comments: "",
  };
  editClient.open = false;
}
function resetAddClient() {
  addClient.title = "";
  addClient.content = {
    companyName: "",
    customerType: "",
    adress: "",
    email: "",
    phone: "",
    managerName: "",
    comments: "",
  }
  addClient.open = false;
}

function copyTextClient() {}

function onFiltered(filteredItems: TableItem<Client>[]) {
  // Trigger pagination to update the number of buttons/pages due to filtering
  totalRows.value = filteredItems.length;
  currentPage.value = 1;
}
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
.client-table {
  overflow-y: auto;
  overflow-x: hidden;

  scrollbar-width: thin;
  scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
}
</style>
