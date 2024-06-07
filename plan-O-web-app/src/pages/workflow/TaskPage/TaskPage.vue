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
      :id="addTask.id"
      v-model="addTask.open"
      :title="addTask.title"
      :hideFooter="true"
      @hide="resetAddTask"
    >
      <BForm class="p-0" @submit="onSubmitAdd(addTask.content)">
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
              v-model="addTask.content.name"
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
              v-model="addTask.content.createdAt"
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
              v-model="addTask.content.deadline"
              type="datetime-local"
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
              v-model="addTask.content.employee"
              :options="employeeList"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Описание -->
        <BFormGroup
          id="add-input-group-5"
          label="Описание"
          label-for="add-input-5"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-5"
              v-model="addTask.content.description"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton type="submit" variant="success" class="mt-3 me-2"
          >Сохранить</BButton
        >
        <BButton variant="outline-danger" class="mt-3" @click="resetAddTask()">
          Отмена</BButton
        >
      </BForm>
    </BModal>

    <!-- EDIT modal -->
    <BModal
      size="lg"
      :id="editTask.id"
      v-model="editTask.open"
      :title="editTask.title"
      :hideFooter="true"
      @hide="resetEditTask"
    >
      <BForm class="p-0" @submit="onSubmitEdit(editTask.content)">
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
              v-model="editTask.content.name"
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
              v-model="editTask.content.createdAt"
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
              v-model="editTask.content.deadline"
              type="datetime-local"
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
              v-model="editTask.content.employee"
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
              v-model="editTask.content.description"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton type="submit" variant="success" class="mt-3 me-2"
          >Сохранить</BButton
        >
        <BButton variant="outline-danger" class="mt-3" @click="resetEditTask()">
          Отмена</BButton
        >
      </BForm>
    </BModal>

    <!-- Info modal -->
    <BModal
      size="lg"
      :id="infoTask.id"
      v-model="infoTask.open"
      :title="infoTask.title"
      :hideFooter="true"
      @hide="resetInfoTask"
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
              v-model="infoTask.content.name"
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
              v-model="infoTask.content.createdAt"
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
              v-model="infoTask.content.deadline"
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
              v-model="infoTask.content.employee"
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
              v-model="infoTask.content.description"
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
import KanbanComponent from "../../../modules/Kanban/KanbanTask.vue";

interface Task {
  id: number;
  name: string;
  deadline: string;
  description: string;
  employee: string;
  createdAt: string;
  status: string; // this will be the status name after transformation
}

const statusList = [
  "На очереди",
  "В работе",
  "Завершен",
];

const employeeList = [
  'Елена Древинг',
  'Алексей Иванов'
]


const itemsTyped = [
  {
    id: 1,
    name: "Task 1",
    deadline: "2024-06-07T12:00:00Z",
    description: "Description for Task 1",
    employee: 'Елена Древинг',
    createdAt: "2024-06-01T08:00:00Z",
    status: 'На очереди',
  },
  {
    id: 2,
    name: "Task 2",
    deadline: "2024-06-08T12:00:00Z",
    description: "Description for Task 2",
    employee: 'Алексей Иванов',
    createdAt: "2024-06-02T08:00:00Z",
    status: "В работе",
  },
];
const convertToIsoWithoutTimezone = (timestamp: number): string => {
  const date = new Date(timestamp * 1000); // Convert Unix timestamp to milliseconds
  date.setHours(date.getHours() + 3); // Add 3 hours to the date
  return date.toISOString().split(".")[0]; // Remove the milliseconds and timezone
};
// Info
const infoTask = reactive({
  open: false,
  id: "info-modal",
  title: "",
  content: {},
});

function info(task: TableItem<Task>, index: number) {
  infoTask.title = `# ${task.name}`;
  infoTask.open = true;
  infoTask.content = { ...task };
}

function resetInfoTask() {
  infoTask.title = "";
  infoTask.content = {
    id: "",
    name: "",
    deadline: "",
    description: "",
    employee: "",
    createdAt: "",
    status: "",
  };
  infoTask.open = false;
}

// Add
const addTask = reactive({
  open: false,
  id: "add-modal",
  title: "",
  content: {},
});

function add() {
  addTask.title = "Новый заказ";
  addTask.open = true;
  addTask.content = {
    id: "",
    name: "",
    deadline: "",
    description: "",
    employee: "",
    createdAt: convertToIsoWithoutTimezone(Date.now() / 1000),
    status: "",
  };
}
function resetAddTask() {
  addTask.title = "";
  addTask.content = {
    id: "",
    name: "",
    deadline: "",
    description: "",
    employee: "",
    createdAt: "",
    status: "",
  };
  addTask.open = false;
}
const onSubmitAdd = (obj) => {
  itemsTyped.push(obj);
  addTask.open = false;
  // TODO:
  //  Добавить метод по отправке даннных на сервер в
};

// Edit
const editTask = reactive({
  open: false,
  id: "edit-modal",
  title: "",
  content: {},
});

function edit(task: TableItem<Task>, index: number) {
  editTask.title = `# ${task.name}`;
  editTask.content = task;
  editTask.open = true;
}

function resetEditTask() {
  editTask.title = "";
  editTask.open = false;
  editTask.content = {};
}
const onSubmitEdit = () => {
  resetEditTask();
  // TODO:
  //  Добавить метод по отправке даннных на сервер в
};
</script>

<style></style>
