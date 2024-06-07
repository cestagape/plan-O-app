<template>
  <BContainer class="p-2 mx-0" :fluid="true">
    <BButtonGroup class="col-1 mb-3 align-items-end">
      <BButton
        :variant="kanbanShow ? 'dark' : 'outline-light'"
        @click="
          () => {
            kanbanShow = true;
            tableShow = false;
          }
        "
      >
        Канбан</BButton
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

    <!-- канбан -->

    <KanbanComponent
      v-if="kanbanShow"
      :columns="statuses"
      :list="itemsTyped"
      :info="info"
      :edit="edit"
      :add="add"
    />

    <div v-if="tableShow">
      <div class="d-flex mb-2" v-if="tableShow">
        <BButton
          type="submit"
          class="text-info"
          variant="dark me-2"
          @click="add('В обработке')"
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
            variant="light"
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
              />
            </svg>
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
              viewBox="0 0 16 16"
            >
              <path
                d="M12.854.146a.5.5 0 0 0-.707 0L10.5 1.793 14.207 5.5l1.647-1.646a.5.5 0 0 0 0-.708zm.646 6.061L9.793 2.5 3.293 9H3.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.207zm-7.468 7.468A.5.5 0 0 1 6 13.5V13h-.5a.5.5 0 0 1-.5-.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.5-.5V10h-.5a.5.5 0 0 1-.175-.032l-.179.178a.5.5 0 0 0-.11.168l-2 5a.5.5 0 0 0 .65.65l5-2a.5.5 0 0 0 .168-.11z"
              />
            </svg>
            Изменить</BButton
          >
        </template>
      </BTable>
    </div>
    <!-- Info modal -->
    <BModal
      size="lg"
      :id="infoOrder.id"
      v-model="infoOrder.open"
      :title="infoOrder.title"
      :ok-only="true"
      @hide="resetInfoOrder"
    >
      <BForm class="p-0">
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

        <!-- createdAt -->
        <BFormGroup
          id="info-input-group-2"
          label="Дата размещения"
          label-for="info-input-2"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-2"
              v-model="infoOrder.content.createdAt"
              type="datetime-local"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyTextOrder(info)" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>

        <!-- Товар -->
        <BTable
          small
          class="rounded-3"
          style="max-width: 50%"
          :items="infoOrder.content.items"
          :fields="itemsFields"
        >
        </BTable>
        <!-- Статус -->
        <BFormGroup
          id="info-input-group-4"
          label="Статус"
          label-for="info-input-4"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-4"
              v-model="infoOrder.content.status"
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

        <!-- Price -->
        <BFormGroup
          id="info-input-group-5"
          label="Сумма заказа"
          label-for="info-input-5"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-5"
              v-model="infoOrder.content.price"
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

        <!-- Ответственный сотрудник -->
        <BFormGroup
          id="info-input-group-6"
          label="Ответственный сотрудник"
          label-for="info-input-6"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-6"
              v-model="infoOrder.content.employee"
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

        <!-- Дэдлайн -->
        <BFormGroup
          id="info-input-group-7"
          label="Дэдлайн"
          label-for="info-input-7"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-7"
              v-model="infoOrder.content.deadline"
              type="datetime-local"
              disabled
            />
            <BInputGroupAppend>
              <BButton @click="copyTextOrder" variant="outline-dark"
                >Copy</BButton
              >
            </BInputGroupAppend>
          </BInputGroup>
        </BFormGroup>

        <!-- Тип оплаты -->
        <BFormGroup
          id="info-input-group-8"
          label="Тип оплаты"
          label-for="info-input-8"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-8"
              v-model="infoOrder.content.paymentType"
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

        <!-- Адрес доставки -->
        <BFormGroup
          id="info-input-group-9"
          label="Адрес доставки"
          label-for="info-input-9"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-9"
              v-model="infoOrder.content.addr"
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

        <!-- Информация о доставке -->
        <BFormGroup
          id="info-input-group-10"
          label="Информация о доставке"
          label-for="info-input-10"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-9"
              v-model="infoOrder.content.deliveryInfo"
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

        <!-- Дополнительная информация -->
        <BFormGroup
          id="info-input-group-11"
          label="Дополнительная информация"
          label-for="info-input-11"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="info-input-9"
              v-model="infoOrder.content.additional"
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

    <!-- Edit modal -->
    <BModal
      size="lg"
      :id="editOrder.id"
      v-model="editOrder.open"
      :title="editOrder.title"
      :hideFooter="true"
      @hide="resetEditOrder"
    >
      <BForm class="p-0" @submit="onSubmitEdit()">
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
              v-model="editOrder.content.companyName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- createdAt -->
        <BFormGroup
          id="edit-input-group-2"
          label="Дата размещения"
          label-for="edit-input-2"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-2"
              v-model.lazy="editOrder.content.createdAt"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Товар -->
        <BTable
          small
          class="rounded-3 mb-1"
          style="max-width: 50%"
          :items="editOrder.content.items"
          :fields="itemsFields"
        >
          <template #cell(name)>
            <div v-for="item in editOrder.content.items" :key="item.id">
              <BFormSelect
                class="m-1 ms-0"
                :options="products"
                v-model="item.name"
              />
            </div>
          </template>

          <template #cell(quantity)>
            <div v-for="item in editOrder.content.items" :key="item.id">
              <BFormInput class="m-1 ms-0" v-model="item.quantity" />
            </div>
          </template>
          <template #cell(actions)="row">
            <BButton
              variant="outline-dark text-danger"
              class="m-1 ms-0"
              @click="removeOrderItem(editOrder.content.items, row.index)"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="1.5em"
                height="1.5em"
                fill="currentColor"
                class="bi bi-x-circle align-top"
                viewBox="0 0 16 16"
              >
                <path
                  d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16"
                />
                <path
                  d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708"
                />
              </svg>
            </BButton>
          </template>
        </BTable>

        <BButton
          block
          @click="addOrderItem(editOrder.content.items)"
          variant="primary"
          class="pb-2 w-50"
          ><svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            class="bi bi-plus-circle-fill p-0 m-0"
            viewBox="0 0 16 16"
          >
            <path
              d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M8.5 4.5a.5.5 0 0 0-1 0v3h-3a.5.5 0 0 0 0 1h3v3a.5.5 0 0 0 1 0v-3h3a.5.5 0 0 0 0-1h-3z"
            />
          </svg>
        </BButton>

        <!-- Статус -->
        <BFormGroup
          id="edit-input-group-4"
          label="Статус"
          label-for="edit-input-4"
          class="my-1"
        >
          <BInputGroup>
            <BFormSelect
              id="edit-input-4"
              v-model="editOrder.content.status"
              :options="statuses"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Price -->
        <BFormGroup
          id="edit-input-group-5"
          label="Сумма заказа"
          label-for="edit-input-5"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-5"
              v-model="editOrder.content.price"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Ответственный сотрудник -->
        <BFormGroup
          id="edit-input-group-6"
          label="Ответственный сотрудник"
          label-for="edit-input-6"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-6"
              v-model="editOrder.content.employee"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Дэдлайн -->
        <BFormGroup
          id="edit-input-group-7"
          label="Дэдлайн"
          label-for="edit-input-7"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-7"
              v-model="editOrder.content.deadline"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Тип оплаты -->
        <BFormGroup
          id="edit-input-group-8"
          label="Тип оплаты"
          label-for="edit-input-8"
          class="my-1"
        >
          <BInputGroup>
            <BFormSelect
              id="edit-input-8"
              v-model="editOrder.content.paymentType"
              :options="paymentTypes"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Адрес доставки -->
        <BFormGroup
          id="edit-input-group-9"
          label="Адрес доставки"
          label-for="edit-input-9"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-9"
              v-model="editOrder.content.addr"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Информация о доставке -->
        <BFormGroup
          id="edit-input-group-10"
          label="Информация о доставке"
          label-for="edit-input-10"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-9"
              v-model="editOrder.content.deliveryInfo"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Дополнительная информация -->
        <BFormGroup
          id="edit-input-group-11"
          label="Дополнительная информация"
          label-for="edit-input-11"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="edit-input-9"
              v-model="editOrder.content.additional"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton
          type="submit"
          variant="success"
          class="mt-3 me-2"
          @click="onSubmitEdit()"
          >Сохранить</BButton
        >
        <BButton
          variant="outline-danger"
          class="mt-3"
          @click="resetEditOrder()"
        >
          Отмена</BButton
        >
      </BForm>
    </BModal>

    <!-- Add modal -->
    <BModal
      size="lg"
      :id="addOrder.id"
      v-model="addOrder.open"
      :title="addOrder.title"
      :hideFooter="true"
      @hide="resetAddOrder"
    >
      <BForm class="p-0" @submit="onSubmitAdd(addOrder.content)">
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
              v-model="addOrder.content.companyName"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- createdAt -->
        <BFormGroup
          id="add-input-group-2"
          label="Дата размещения"
          label-for="add-input-2"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-2"
              v-model.lazy="addOrder.content.createdAt"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Товар -->
        <!-- дописать v-model -->
        <BTable
          small
          class="rounded-3 mb-1"
          style="max-width: 50%"
          :items="addOrder.content.items"
          :fields="itemsFields"
        >
          <template #cell(name)>
            <div v-for="item in addOrder.content.items" :key="item.id">
              <BFormSelect
                class="m-1 ms-0"
                :options="products"
                v-model="item.name"
              />
            </div>
          </template>

          <template #cell(quantity)>
            <div v-for="item in addOrder.content.items" :key="item.id">
              <BFormInput class="m-1 ms-0" v-model="item.quantity" />
            </div>
          </template>
          <template #cell(actions)="row">
            <BButton
              variant="outline-dark text-danger"
              class="m-1 ms-0"
              @click="removeOrderItem(addOrder.content.items, row.index)"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="1.5em"
                height="1.5em"
                fill="currentColor"
                class="bi bi-x-circle align-top"
                viewBox="0 0 16 16"
              >
                <path
                  d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16"
                />
                <path
                  d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708"
                />
              </svg>
            </BButton>
          </template>
        </BTable>
        <BButton
          block
          @click="addOrderItem(addOrder.content.items)"
          variant="dark"
          class="pb-2 w-50"
          ><svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            class="bi bi-plus-circle-fill p-0 m-0"
            viewBox="0 0 16 16"
          >
            <path
              d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M8.5 4.5a.5.5 0 0 0-1 0v3h-3a.5.5 0 0 0 0 1h3v3a.5.5 0 0 0 1 0v-3h3a.5.5 0 0 0 0-1h-3z"
            />
          </svg>
        </BButton>
        <!-- Статус -->
        <BFormGroup
          id="edit-input-group-4"
          label="Статус"
          label-for="edit-input-4"
          class="my-1"
        >
          <BInputGroup>
            <BFormSelect
              id="edit-input-4"
              v-model="editOrder.content.status"
              :options="statuses"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Price -->
        <BFormGroup
          id="add-input-group-5"
          label="Сумма заказа"
          label-for="add-input-5"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-5"
              v-model="addOrder.content.price"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Ответственный сотрудник -->
        <BFormGroup
          id="add-input-group-6"
          label="Ответственный сотрудник"
          label-for="add-input-6"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-6"
              v-model="addOrder.content.employee"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Дэдлайн -->
        <BFormGroup
          id="add-input-group-7"
          label="Дэдлайн"
          label-for="add-input-7"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-7"
              v-model="addOrder.content.deadline"
              type="datetime-local"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Тип оплаты -->
        <BFormGroup
          id="add-input-group-8"
          label="Тип оплаты"
          label-for="add-input-8"
          class="my-1"
        >
          <BInputGroup>
            <BFormSelect
              id="add-input-8"
              v-model="addOrder.content.paymentType"
              :options="paymentTypes"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Адрес доставки -->
        <BFormGroup
          id="add-input-group-9"
          label="Адрес доставки"
          label-for="add-input-9"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-9"
              v-model="addOrder.content.addr"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Информация о доставке -->
        <BFormGroup
          id="add-input-group-10"
          label="Информация о доставке"
          label-for="add-input-10"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-9"
              v-model="addOrder.content.deliveryInfo"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>

        <!-- Дополнительная информация -->
        <BFormGroup
          id="add-input-group-11"
          label="Дополнительная информация"
          label-for="add-input-11"
          class="my-1"
        >
          <BInputGroup>
            <BFormInput
              id="add-input-9"
              v-model="addOrder.content.additional"
              type="text"
            />
          </BInputGroup>
        </BFormGroup>
        <BButton type="submit" variant="success" class="mt-3 me-2"
          >Сохранить</BButton
        >
        <BButton variant="outline-danger" class="mt-3" @click="resetAddOrder()">
          Отмена</BButton
        >
      </BForm>
    </BModal>
  </BContainer>
</template>
<script setup lang="ts">
import KanbanComponent from "../../../modules/Kanban/KanbanCRM.vue";

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
import { reactive, ref } from "vue";

const kanbanShow = ref(true);
const tableShow = ref(false);

interface item {
  id: number;
  name: string;
  price: number;
  quantity: number;
}

interface Order {
  id: number;
  companyName: string;
  createdAt: string; // ISO 8601 date string for TIMESTAMPTZ
  design: string;
  status: string;
  price: number;
  addr: string;
  employee: number;
  deadline: string; // ISO 8601 date string for TIMESTAMPTZ
  paymentType: string;
  deliveryInfo: Record<string, any>; // JSONB field
  additional: string;
  items: item[];
}
const employees = [];

const statuses = [
  "В процессе",
  "Ожидает оплаты",
  "В производстве",
  "Контроль качества",
  "Упаковка",
  "Документация",
  "Отправлено в доставку",
];

const products = ["AWS-пластик", "Наклейка", "Шильд", "Шелкография"];

const paymentTypes = ["Рассчетный счет компании", "Наличные", "СБП"];

const itemsFields: Exclude<TableFieldRaw<item>, string>[] = [
  {
    key: "name",
    label: "Наименование",
  },
  {
    key: "quantity",
    label: "Количество",
  },
  {
    key: "actions",
    label: "",
  },
];

const fieldsTyped: Exclude<TableFieldRaw<Order>, string>[] = [
  {
    key: "id",
    label: "ID заказа",
    sortable: true,
    sortDirection: "asc",
  },
  {
    key: "companyName",
    label: "Клиент",
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
    key: "items",
    label: "Товары",
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

// Info
const infoOrder = reactive({
  open: false,
  id: "info-modal",
  title: "",
  content: {},
});

function info(order: TableItem<Order>, index: number) {
  infoOrder.title = `# ${order.id}`;
  infoOrder.open = true;
  infoOrder.content = { ...order };
}

function resetInfoOrder() {
  infoOrder.title = "";
  infoOrder.content = {
    id: "",
    companyName: "",
    createdAt: "",
    design: "",
    status: "",
    price: "",
    addr: "",
    employee: "",
    deadline: "",
    paymentType: "",
    deliveryInfo: "",
    additional: "",
    items: [],
  };
  infoOrder.open = false;
}

// Edit
const editOrder = reactive({
  open: false,
  id: "edit-modal",
  title: "",
  content: {},
});

function edit(order: TableItem<Order>, index: number) {
  editOrder.title = `# ${order.id}`;
  editOrder.content = order;
  editOrder.open = true;
}

function resetEditOrder() {
  editOrder.title = "";
  editOrder.open = false;
  editOrder.content = {};
}
const onSubmitEdit = () => {
  // alert(JSON.stringify(item))
  editOrder.open = false;
  // TODO:
  //  Добавить метод по отправке даннных на сервер в
};
// Add
const addOrder = reactive({
  open: false,
  id: "add-modal",
  title: "",
  content: {},
});

function add(status) {
  addOrder.title = "Новый заказ";
  addOrder.open = true;
  addOrder.content = {
    id: "",
    companyName: "",
    createdAt: convertToIsoWithoutTimezone(Date.now() / 1000),
    design: "",
    status: status,
    price: "",
    addr: "",
    employee: "",
    deadline: "",
    paymentType: "",
    deliveryInfo: "",
    additional: "",
    items: [],
  };
}
function addOrderItem(obj) {
  obj.push({
    id: obj.length + 1,
    name: "",
    price: 0,
    quantity: 0,
  });
}
function removeOrderItem(arr: Array<item>, index: number) {
  arr.splice(index, 1);
}
function resetAddOrder() {
  addOrder.title = "";
  addOrder.content = {
    id: "",
    companyName: "",
    createdAt: "",
    design: "",
    status: "",
    price: "",
    addr: "",
    employee: "",
    deadline: "",
    paymentType: "",
    deliveryInfo: "",
    additional: "",
    items: [],
  };
  addOrder.open = false;
}
const onSubmitAdd = (obj) => {
  itemsTyped.push(obj);
  addOrder.open = false;
  // TODO:
  //  Добавить метод по отправке даннных на сервер в
};

const convertToIsoWithoutTimezone = (timestamp: number): string => {
  const date = new Date(timestamp * 1000); // Convert Unix timestamp to milliseconds
  date.setHours(date.getHours() + 3); // Add 3 hours to the date
  return date.toISOString().split(".")[0]; // Remove the milliseconds and timezone
};


const itemsTyped: Order[] = reactive([
  {
    id: 1,
    companyName: "Black Co",
    createdAt: convertToIsoWithoutTimezone(Date.now() / 1000),
    design: "Design A",
    status: statuses[0], // inProcess
    price: 1000,
    addr: "123 Main St",
    employee: 101,
    deadline: convertToIsoWithoutTimezone(Date.now() / 1000 + 7 * 24 * 60 * 60), // 1 week from now
    paymentType: paymentTypes[0], // рассчетный счет компании
    deliveryInfo: { courier: "DHL", trackingNumber: "123456789" },
    additional: "Handle with care",
    items: [
      {
        id: 1,
        name: "Наклейка",
        quantity: 5,
        price: 200,
      },
    ],
  },

  {
    id: 2,
    companyName: "White Co",
    createdAt: convertToIsoWithoutTimezone(Date.now() / 1000),
    design: "Design B",

    status: statuses[1], // paymentPending
    price: 500,
    addr: "456 Elm St",
    employee: 102,
    deadline: convertToIsoWithoutTimezone(
      Date.now() / 1000 + 14 * 24 * 60 * 60
    ), // 2 weeks from now
    paymentType: paymentTypes[1], // наличные
    deliveryInfo: { courier: "UPS", trackingNumber: "987654321" },
    additional: "Fragile item",
    items: [
      {
        id: 2,
        name: "AWS-пластик",
        quantity: 10,
        price: 50,
      },
    ],
  },
  {
    id: 3,
    companyName: "Green Co",
    createdAt: convertToIsoWithoutTimezone(1717425145),
    design: "Design C",
    status: statuses[2], // inProduction
    price: 2000,
    addr: "789 Oak St",
    employee: 103,
    deadline: convertToIsoWithoutTimezone(1717435945), // 3 days from now
    paymentType: paymentTypes[2], // сбп
    deliveryInfo: { courier: "FedEx", trackingNumber: "1122334455" },
    additional: "Urgent delivery",
    items: [
      {
        id: 3,
        name: "AWS-пластик",
        quantity: 10,
        price: 200,
      },
    ],
  },
  {
    id: 4,
    companyName: "Yellow Co",
    createdAt: convertToIsoWithoutTimezone(Date.now() / 1000),
    design: "Design D",
    status: statuses[3], // qualityControl
    price: 1500,
    addr: "101 Maple St",
    employee: 104,
    deadline: convertToIsoWithoutTimezone(
      Date.now() / 1000 + 10 * 24 * 60 * 60
    ), // 10 days from now
    paymentType: paymentTypes[0], // рассчетный счет компании
    deliveryInfo: { courier: "DHL", trackingNumber: "2233445566" },
    additional: "Inspect carefully",
    items: [
      {
        id: 4,
        name: "AWS-пластик",
        quantity: 10,
        price: 150,
      },
      { id: 10, name: "Наклейка", quantity: 220, price: 150 },
    ],
  },
  {
    id: 5,
    companyName: "Brown Co",
    createdAt: convertToIsoWithoutTimezone(Date.now() / 1000),
    design: "Design E",
    status: statuses[4], // packaging
    price: 2500,
    addr: "202 Birch St",
    employee: 105,
    deadline: convertToIsoWithoutTimezone(Date.now() / 1000 + 5 * 24 * 60 * 60), // 5 days from now
    paymentType: paymentTypes[1], // наличные
    deliveryInfo: { courier: "UPS", trackingNumber: "5566778899" },
    additional: "Package securely",
    items: [
      {
        id: 5,
        name: "AWS-пластик",
        quantity: 10,
        price: 250,
      },
    ],
  },
]);

const totalRows = ref(itemsTyped.length);
const currentPage = ref(1);
const perPage = ref(itemsTyped.length);
const filter = ref("");
const filterOn = ref([]);

function onFiltered(filteredItems: TableItem<Order>[]) {
  // Trigger pagination to update the number of buttons/pages due to filtering
  totalRows.value = filteredItems.length;
  currentPage.value = 1;
}
</script>
<style>
.order-table {
  overflow-y: auto;
  overflow-x: hidden;

  scrollbar-width: thin;
  scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
}
</style>
