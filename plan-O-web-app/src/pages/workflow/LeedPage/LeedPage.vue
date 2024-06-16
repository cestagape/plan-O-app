<template>
  <BContainer class="p-2 px-2 mx-0" :fluid="true">
    <KanbanComponent
      :columns="statusList"
      :list="itemsTyped"
      :info="info"
      :edit="edit"
      :add="add"
    />
    <!-- Add modal -->
    <BModal
      size="lg"
      :id="addLead.id"
      v-model="addLead.open"
      :title="addLead.title"
      :hideFooter="true"
      @hide="resetAddLead"
    >
      <BForm class="p-0" @submit="onSubmitAdd(addLead.content)">
        <!-- Наименование задачи -->
        <BFormGroup
          id="add-input-group-1"
          label="Наименование задачи"
          label-for="add-input-1"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-1"
              v-model="addLead.content.name"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Размещен -->
        <BFormGroup
          id="add-input-group-2"
          label="Размещен"
          label-for="add-input-2"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-2"
              v-model="addLead.content.createdAt"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Дэдлайн -->
        <BFormGroup
          id="add-input-group-3"
          label="Дэдлайн"
          label-for="add-input-3"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-3"
              v-model="addLead.content.deadline"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!--  название компании -->
        <BFormGroup
          id="add-input-group-4"
          label="Название компании"
          label-for="add-input-4"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-4"
              v-model="addLead.content.companyName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Ответственный -->
        <BFormGroup
          id="edit-input-group-5"
          label="Ответственный"
          label-for="edit-input-5"
          class="my-1"
        >
          <BInputGroup>
            <BFormSelect
              id="edit-input-5"
              v-model="addLead.content.employee"
              :options="employeeList"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Описание -->
        <BFormGroup
          id="add-input-group-6"
          label="Описание"
          label-for="add-input-6"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-6"
              v-model="addLead.content.description"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton type="submit" variant="success" class="mt-3 me-2"
          >Сохранить</BButton
        >
        <BButton variant="outline-danger" class="mt-3" @click="resetAddLead()">
          Отмена</BButton
        >
      </BForm>
    </BModal>

    <!-- EDIT modal -->
    <BModal
      size="lg"
      :id="editLead.id"
      v-model="editLead.open"
      :title="editLead.title"
      :hideFooter="true"
      @hide="resetEditLead"
    >
      <BForm class="p-0" @submit="onSubmitEdit(editLead.content)">
        <!-- Наименование задачи -->
        <BFormGroup
          id="edit-input-group-1"
          label="Наименование задачи"
          label-for="edit-input-1"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-1"
              v-model="editLead.content.name"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Размещен -->
        <BFormGroup
          id="edit-input-group-2"
          label="Размещен"
          label-for="edit-input-2"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-2"
              v-model="editLead.content.createdAt"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Дэдлайн -->
        <BFormGroup
          id="edit-input-group-3"
          label="Дэдлайн"
          label-for="edit-input-3"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-3"
              v-model="editLead.content.deadline"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!--  название компании -->
        <BFormGroup
          id="add-input-group-4"
          label="Название компании"
          label-for="add-input-4"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-4"
              v-model="addLead.content.companyName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <!-- Ответственный -->
        <BFormGroup
          id="edit-input-group-4"
          label="Ответственный"
          label-for="edit-input-4"
          class="my-1"
        >
          <BInputGroup>
            <BFormSelect
              id="edit-input-4"
              v-model="editLead.content.employee"
              :options="employeeList"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Описание -->
        <BFormGroup
          id="edit-input-group-5"
          label="Описание"
          label-for="edit-input-5"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-5"
              v-model="editLead.content.description"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton type="submit" variant="success" class="mt-3 me-2"
          >Сохранить</BButton
        >
        <BButton variant="outline-danger" class="mt-3" @click="resetEditLead()">
          Отмена</BButton
        >
      </BForm>
    </BModal>

    <!-- Info modal -->
    <BModal
      size="lg"
      :id="infoLead.id"
      v-model="infoLead.open"
      :title="infoLead.title"
      :hideFooter="true"
      @hide="resetInfoLead"
    >
      <BForm class="p-0">
        <!-- Наименование задачи -->
        <BFormGroup
          id="info-input-group-1"
          label="Наименование задачи"
          label-for="info-input-1"
          class="mb-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-1"
              v-model="infoLead.content.name"
              type="text"
              disabled
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Размещен -->
        <BFormGroup
          id="info-input-group-2"
          label="Размещен"
          label-for="info-input-2"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-2"
              v-model="infoLead.content.createdAt"
              type="datetime-local"
              disabled
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Дэдлайн -->
        <BFormGroup
          id="info-input-group-3"
          label="Дэдлайн"
          label-for="info-input-3"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-3"
              v-model="infoLead.content.deadline"
              type="datetime-local"
              disabled
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Ответственный -->
        <BFormGroup
          id="-input-group-4"
          label="Ответственный"
          label-for="info-input-4"
          class="my-1"
        >
          <BInputGroup>
            <BFormSelect
              id="info-input-4"
              v-model="infoLead.content.employee"
              :options="employeeValues"
              disabled
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Описание -->
        <BFormGroup
          id="info-input-group-5"
          label="Описание"
          label-for="info-input-5"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-5"
              v-model="infoLead.content.description"
              type="text"
              disabled
            />
          </BInputGroup>
        </BFormGroup>
      </BForm>
    </BModal>
  </BContainer>
</template>

<script setup lang="ts">
import {
  BButton,
  BInputGroup,
  BFormGroup,
  BFormInput,
  BInputGroupAppend,
  BModal,
  BContainer,
  BTable,
  type BTableSortBy,
  type TableFieldRaw,
  type TableItem,
} from "bootstrap-vue-next";
import { reactive } from "vue";
import KanbanComponent from "../../../modules/Kanban/KanbanLead.vue";

interface Lead {
  id: number;
  name: string;
  deadline: string;
  description: string;
  employee: string;
  createdAt: string;
  companyName: string;
  status: string; // this will be the status name after transformation
}

const statusList = ["На очереди", "В работе", "Завершен"];

const employeeList = ["Елена Древинг", "Алексей Иванов"];

const convertToIsoWithoutTimezone = (timestamp: number): string => {
    const date = new Date(timestamp);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are zero-based
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    const formattedDateTime = `${year}-${month}-${day}T${hours}:${minutes}`;

    return formattedDateTime;
}

const itemsTyped = [
  {
    id: 1,
    name: "Lead 1",
    deadline: "2024-06-07T12:00:00Z",
    description: "Description for Lead 1",
    companyName: "Company 1",
    employee: "Елена Древинг",
    createdAt: "2024-06-01T08:00:00Z",
    status: "На очереди",
  },
  {
    id: 2,
    name: "Lead 2",
    deadline: "2024-06-08T12:00:00Z",
    description: "Description for Lead 2",
    companyName: "Company 2",
    employee: "Алексей Иванов",
    createdAt: "2024-06-02T08:00:00Z",
    status: "В работе",
  },
];

// Info
const infoLead = reactive({
  open: false,
  id: "info-modal",
  title: "",
  content: {},
});

function info(lead: TableItem<Lead>, index: number) {
  infoLead.title = `# ${lead.name}`;
  infoLead.open = true;
  infoLead.content = { ...lead };
}

function resetInfoLead() {
  infoLead.title = "";
  infoLead.content = {} as Lead;
  infoLead.open = false;
}

// Add
const addLead = reactive({
  open: false,
  id: "add-modal",
  title: "",
  content: {},
});

function add() {
  addLead.title = "Новый лид";
  addLead.open = true;
  addLead.content = {
    id: "",
    name: "",
    deadline: "",
    description: "",
    employee: "",
    createdAt: convertToIsoWithoutTimezone(Date.now() / 1000),
    status: "",
  };
}
function resetAddLead() {
  addLead.title = "";
  addLead.content = {
    id: "",
    name: "",
    deadline: "",
    description: "",
    employee: "",
    createdAt: "",
    status: "",
  };
  addLead.open = false;
}
const onSubmitAdd = (obj) => {
  itemsTyped.push(obj);
  addLead.open = false;
  // TODO:
  //  Добавить метод по отправке даннных на сервер в
};

// Edit
const editLead = reactive({
  open: false,
  id: "edit-modal",
  title: "",
  content: {},
});

function edit(lead: TableItem<Lead>, index: number) {
  editLead.title = `# ${lead.name}`;
  editLead.content = lead;
  editLead.open = true;
}

function resetEditLead() {
  editLead.title = "";
  editLead.open = false;
  editLead.content = {};
}
const onSubmitEdit = () => {
  resetEditLead();
  // TODO:
  //  Добавить метод по отправке даннных на сервер в
};
</script>

<style></style>
