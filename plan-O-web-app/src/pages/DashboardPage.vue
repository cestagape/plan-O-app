<template>
  <BContainer class="p-2 mx-0" :fluid="true">
    <BCardGroup deck class="pb-2">
      <BCard
        header="Среднее время жизни заказа"
        header-tag="header"
        header-bg-variant="dark"
        body-bg-variant="dark py-3 bg-opacity-75 rounded-bottom"
        class="shadow text-light px-0 bg-transparent"
      >
        <div
          class="d-flex justify-content-between align-items-center w-100 px-3"
        >
          <p
            class="fs-4 fw-medium text-center rounded-2 px-2 w-50 h-50 bg-dark m-0"
          >
            {{ avgTime }} часа
          </p>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="5em"
            height="5em"
            fill="black"
            class="bi bi-clock p-1 rounded-2"
            viewBox="-4 -4 24 24"
            style="background-color: #b6f3ff"
          >
            <path
              d="M8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71z"
            />
            <path
              d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m7-8A7 7 0 1 1 1 8a7 7 0 0 1 14 0"
            />
          </svg>
        </div>
      </BCard>

      <BCard
        header="Средний чек"
        header-tag="header"
        header-bg-variant="dark"
        body-bg-variant="dark py-3 bg-opacity-75 rounded-bottom"
        class="shadow text-light px-0 bg-transparent"
      >
        <div
          class="d-flex justify-content-between align-items-center w-100 px-3"
        >
          <p
            class="fs-4 fw-medium text-center rounded-2 px-2 w-50 h-100 bg-dark m-0"
          >
            {{ avgPayCheck }} рублей
          </p>
          <div class="col-1"></div>
          <svg
            width="5em"
            height="5em"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            class="bi bi-clock p-1 rounded-2"
            style="background-color: #b6f3ff"
          >
            <path
              d="M11 8V17"
              stroke="#323232"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M9 15H15"
              stroke="#323232"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M11 8H13.5C16.5 8 16.5 12 13.5 12H9"
              stroke="#323232"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </div>
      </BCard>

      <BCard
        header="Выручка за месяц"
        header-tag="header"
        header-bg-variant="dark"
        body-bg-variant="dark py-3 bg-opacity-75 rounded-bottom"
        class="shadow text-light px-0 bg-transparent"
      >
        <div
          class="d-flex justify-content-between align-items-center w-100 px-3"
        >
          <p
            class="fs-4 fw-medium text-center rounded-2 px-2 w-50 h-50 bg-dark m-0"
          >
            {{ revenue }} рублей
          </p>
          <div class="col-1"></div>
          <svg
            width="5em"
            height="5em"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            class="bi bi-clock p-1 rounded-2"
            style="background-color: #b6f3ff"
          >
            <path
              d="M11 8V17"
              stroke="#323232"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M9 15H15"
              stroke="#323232"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M11 8H13.5C16.5 8 16.5 12 13.5 12H9"
              stroke="#323232"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </div>
      </BCard>
    </BCardGroup>
    <BCardGroup deck>
      <BCard
        header="Среднее выполнения время заказа по месяцам"
        header-tag="header"
        header-bg-variant="dark"
        body-bg-variant="dark bg-opacity-75 py-3  rounded-bottom"
        class="shadow text-light px-0 bg-transparent"
      >
        <canvas ref="avgOrderTime"></canvas>
      </BCard>
      <BCard
        header="Среднее время производства заказа"
        header-tag="header"
        header-bg-variant="dark"
        body-bg-variant="dark bg-opacity-75 py-3  rounded-bottom"
        class="shadow text-light px-0 bg-transparent"
      >
        <canvas ref="avgProdTime"></canvas>
      </BCard>
    </BCardGroup>
  </BContainer>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Chart from "chart.js/auto";


const avgTime = 144;
const revenue = '101,685';
const avgPayCheck = '30,000';

const months = [
  "Январь",   // January
  "Февраль",  // February
  "Март",     // March
  "Апрель",   // April
  "Май",      // May
  "Июнь",     // June
  "Июль",     // July
  "Август",   // August
  "Сентябрь", // September
  "Октябрь",  // October
  "Ноябрь",   // November
  "Декабрь"   // December
];

Chart.defaults.color = '#FFF';

const avgOrderTime = ref(null);
const avgProdTime = ref(null);


onMounted(() => {
  if (avgOrderTime.value) {
    new Chart(avgOrderTime.value, {
      type: "line",
      data: {
        labels: months,
        datasets: [
          {
            label: "# of Votes",
            data: [12, 19, 3, 5, 2, 3],
            borderWidth: 1,
          },
        ],
      },
      options: {
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      },
    });
  }

  if (avgProdTime.value) {
    new Chart(avgProdTime.value, {
      type: "line",
      data: {
        labels: months,
        datasets: [
          {
            label: "# of Votes",
            data: [12, 19, 3, 5, 2, 3],
            borderWidth: 1,
          },
        ],
      },
      options: {
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      },
    });
  }
}); // Add a semicolon here
</script>

<style>
.card-deck {
  gap: 0.5em;
}
</style>
