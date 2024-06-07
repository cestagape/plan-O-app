<template>
    <BCardGroup deck>
      <BCard
        v-for="column in columns"
        :key="column"
        :header="column"
        header-tag="header"
        header-text-variant="white"
        header-bg-variant="dark text-center"
        bg-variant="light bg-opacity-50"
        class="px-0 m-0 border-0 shadow-lg cardin"
        bodyClass="p-2 text-center"
      >
        <BButton pill variant="outline-dark" class="mb-3" @click="add(column)">
          Добавить
        </BButton>
        <draggable
          class="list-group"
          :list="getListByColumn(column)"
          group="people"
          itemKey="id"
        >
          <template #item="{ element, index }">
            <BCard
              class="mb-3 text-start m-0 border-dark bg-light bg-opacity-75 shadow"
            >
              <BCardHeader
                class="fs-6 fw-bold p-1 text-center bg-transparent border-dark"
              >
                Лид №{{ element.id }}
              </BCardHeader>
              <BCardBody class="p-1 fw-light">
                <p>
                  Дэдлайн <br /><span class="fw-semibold">{{
                    element.deadline
                  }}</span>
                </p>
                <p>
                  Клиент <br /><span class="fw-semibold">{{
                    element.companyName
                  }}</span>
                </p>
                <p>
                  Ответственный <br />
                  <span class="fw-semibold">{{ element.employee }}</span>
                </p>
                <p>
                  Описание <br /><span class="fw-semibold">{{ element.description }}</span>
                </p>
                <BForm class="text-center">
                  <BButton
                    class="me-1 my-1 fw-light w-100"
                    variant="outline-dark"
                    @click="info(element, index)"
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
                    class="me-1 my-1 fw-light w-100"
                    variant="outline-dark"
                    @click="edit(element, index)"
                    ><svg
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
                </BForm>
              </BCardBody>
            </BCard>
          </template>
        </draggable>
      </BCard>
    </BCardGroup>
  </template>
  
  <script>
  import { BButton, BCardBody, BForm } from "bootstrap-vue-next";
  import draggable from "vuedraggable";
  
  export default {
    components: {
      draggable,
    },
    props: {
      columns: {
        type: Array,
        required: true,
      },
      add: {
        type: Function,
        required: true,
      },
      list: {
        type: Array,
        required: true,
      },
      info: {
        type: Function,
        required: true,
      },
      edit: {
        type: Function,
        required: true,
      },
    },
    data() {
      return {
        lists: this.initializeLists(),
      };
    },
    methods: {
      initializeLists() {
        const lists = {};
        this.columns.forEach((column) => {
          lists[column] = this.list.filter((item) => item.status === column);
        });
        return lists;
      },
      getListByColumn(column) {
        return this.lists[column];
      },
      updateLists() {
        this.lists = this.initializeLists();
      },
    },
  };
  </script>
  
  <style>
  html,
  body,
  #app {
    height: 100%;
    margin: 0;
  }
  
  .BCardGroup {
    height: 100%;
  }
  
  
  .cardin {
    overflow-y: auto;
    overflow-x: hidden;
    scrollbar-width: thin;
    scrollbar-color: rgba(107, 113, 118, 1) rgba(255, 255, 255, 0);
    scroll-margin: 0;
    scroll-padding: 0;
    scrollbar-gutter: 0;
    max-height: calc(100vh - 13em);
    min-height: calc(100vh - 13em);

  }
  
  .card-body {
    padding: 0;
  }
  
  .card-deck {
    gap: 0.5em;
  }
  .list-group {
    height: 100%;
  }
  </style>
  