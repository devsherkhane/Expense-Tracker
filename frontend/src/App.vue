<template>
  <div class="container">

    <h1>Expense Tracker</h1>

    <input v-model="title" placeholder="Title" />
    <input v-model="amount" type="number" placeholder="Amount" />
    <input v-model="category" placeholder="Category" />
    <input v-model="date" type="date" />

    <button @click="addExpense">Add</button>

    <ul>
      <li v-for="e in expenses" :key="e.id">
        {{ e.title }} - ₹{{ e.amount }} ({{ e.category }})
        <button @click="deleteExpense(e.id)">X</button>
      </li>
    </ul>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const expenses = ref([])

const title = ref('')
const amount = ref('')
const category = ref('')
const date = ref('')

const API = "http://localhost:8080/expenses"

const loadExpenses = async () => {
  const res = await axios.get(API)
  expenses.value = res.data
}

const addExpense = async () => {
  await axios.post(API, {
    title: title.value,
    amount: amount.value,
    category: category.value,
    date: date.value
  })

  loadExpenses()
}

const deleteExpense = async (id) => {
  await axios.delete(`${API}/${id}`)
  loadExpenses()
}

onMounted(loadExpenses)
</script>

<style>
.container {
  width: 400px;
  margin: auto;
}
input {
  display: block;
  margin: 5px;
}
</style>
