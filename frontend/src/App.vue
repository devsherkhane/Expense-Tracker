<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const expenses = ref([])
const form = ref({ 
  title: '', 
  amount: null, 
  category: '', 
  date: new Date().toISOString().substr(0, 10) 
})

const API = "http://localhost:8080/expenses"

const totalAmount = computed(() => {
  return expenses.value.reduce((sum, e) => sum + (e.amount || 0), 0)
})

const loadExpenses = async () => {
  try {
    const res = await axios.get(API)
    expenses.value = res.data || []
  } catch (err) {
    console.error("Load failed", err)
  }
}

const handleAdd = async () => {
  if (!form.value.title || !form.value.amount || !form.value.category) return
  try {
    await axios.post(API, form.value)
    form.value = { title: '', amount: null, category: '', date: new Date().toISOString().substr(0, 10) }
    await loadExpenses()
  } catch (err) {
    console.error("Add failed", err)
  }
}

const deleteEntry = async (id) => {
  if (confirm("Remove this item?")) {
    try {
      await axios.delete(`${API}/${id}`)
      await loadExpenses()
    } catch (err) {
      console.error("Delete failed", err)
    }
  }
}

onMounted(loadExpenses)
</script>

<template>
  <div id="app">
    <h1>Expense Tracker</h1>
    
    <div class="container">
      <div style="display: flex; justify-content: space-between; margin-bottom: 1.5rem; align-items: baseline;">
        <span style="font-weight: 700; opacity: 0.7; font-size: 0.9rem; text-transform: uppercase;">Total Spent</span>
        <span style="font-size: 1.5rem; font-weight: 800; color: var(--primary);">₹{{ totalAmount.toFixed(2) }}</span>
      </div>

      <div class="form-group">
        <input v-model="form.title" type="text" placeholder="What did you buy?" />
        
        <div style="display: flex; gap: 10px;">
          <input v-model.number="form.amount" type="number" placeholder="Amount" />
          <select v-model="form.category">
            <option value="" disabled>Category</option>
            <option>Food</option>
            <option>Travel</option>
            <option>Bills</option>
            <option>Shopping</option>
          </select>
        </div>
        
        <button @click="handleAdd" style="width: 100%;">
          Add Transaction
        </button>
      </div>

      <ul>
        <li v-for="e in expenses" :key="e.id" >
          <div style="display: flex; flex-direction: column;">
            <span style="font-weight: 700; font-size: 1rem; color: black;">{{ e.title }}</span>
            <span style="font-size: 0.75rem; color: #94a3b8; font-family: monospace; color: black;">{{ e.date }} • {{ e.category }}</span>
          </div>
          
          <div style="display: flex; align-items: center; gap: 1rem;">
            <span style="font-weight: 700; color: var(--text-main); color: black;">₹{{ e.amount }}</span>
            <button @click="deleteEntry(e.id)" class="delete-btn">
              Delete
            </button>
          </div>
        </li>
      </ul>

      <div v-if="expenses.length === 0" style="text-align: center; padding: 2rem; color: #94a3b8; font-style: italic;">
        No records yet.
      </div>
    </div>
  </div>
</template>

<style scoped>
.form-group {
  margin-bottom: 2rem;
  padding-bottom: 2rem;
  border-bottom: 2px dashed #f1f5f9;
}
</style>