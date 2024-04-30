<script setup>
import { computed, ref, onMounted, watch } from 'vue';
import axios from 'axios'
import EditorPopup from '@/components/EditorPopup.vue'

const state = ref({
  current: {
    id: null,
    term: null,
    description: null
  },
  keyword: '',
  items: [],
  showPopup: false
})

const filteredItems = computed(() => {
  if (state.value.keyword == null || state.value.keyword === '') return state.value.items
  return state.value.items.filter(x => x.term.toLowerCase().includes(state.value.keyword.toLowerCase()))
})

const getList = function() {
  axios.get(import.meta.env.VITE_API_HOST + '/list')
    .then(res => {
      state.value.items = res.data.data
    })
    .catch(err => {
      console.log(err)
    })
}

const selectItem = function(id) {
  state.value.current.id = id
  let item = state.value.items.find(x => x.id === id)
  state.value.current.term = item.term
  state.value.current.description = item.description
}

const add = function() {
  state.value.current.id = null
  state.value.current.term = ''
  state.value.current.description = ''
  state.value.showPopup = true
}

const edit = function() {
  state.value.showPopup = true
}

const save = function(payload) {
  let data = {
    id: payload.id,
    term: payload.term,
    description: payload.description,
  }

  if (payload.action === 'update') {
    axios.put(import.meta.env.VITE_API_HOST + '/update', data)
      .then(res => {})
      .catch(err => console.log(err))
      .finally(() => getList())
    state.value.showPopup = false
    return
  }

  axios.post(import.meta.env.VITE_API_HOST + '/create', data)
      .then(res => {})
      .catch(err => console.log(err))
      .finally(() => getList())
  state.value.showPopup = false
}

const deleteItem = function() {
  axios.delete(import.meta.env.VITE_API_HOST + '/delete/' + state.value.current.id)
      .then(res => {})
      .catch(err => console.log(err))
      .finally(() => getList())
}

onMounted(() => {
  getList()
})

watch(filteredItems, newVal => {
  state.value.current.id = newVal[0].id
  state.value.current.term = newVal[0].term
  state.value.current.description = newVal[0].description
})
</script>

<template>
  <main class="font-raleway p-8">
    <h1 class="text-center mb-8 text-2xl cursor-pointer" @click="add">My Dictionary</h1>
    <div class="flex flex-col w-[1024px] m-auto border-2 rounded-xl">

      <div class="headers w-full flex flex-row">
        <div class="w-[270px] px-4 py-2">
          <input type="text" class="border-[1px] rounded-[24px] text-lg px-4 py-2"
            placeholder="Type to search.." v-model="state.keyword" />
        </div>
        <div class="grow px-4 py-2 flex items-center" >
          Description
          <span v-if="state.current.id != null" class="cursor-pointer" @click="edit"> - Edit</span>
          <span v-if="state.current.id != null" class="cursor-pointer" @click="deleteItem"> - Delete</span>
        </div>
      </div>

      <div class="body w-full flex flex-row h-[500px]">
        <div class="w-[270px] grow-0 shrink-0 basis-[270px] px-4 py-2 overflow-y-auto h-full">
          <ul>
            <li v-for="item in filteredItems" class="item px-2 py-1 hover:bg-[#f2f2f2] cursor-pointer" :class="{ active: item.id === state.current.id }" @click="selectItem(item.id)">{{ item.term }}</li>
          </ul>
        </div>
        <div class="shrink p-4 overflow-y-auto h-full">
          <div>
            {{ state.current.description }}
          </div>
        </div>
      </div>

    </div>

    <EditorPopup v-if="state.showPopup" 
      :id="state.current.id"
      :term="state.current.term"
      :description="state.current.description"
      @onSave="save"
      @onCancel="state.showPopup = false" />
  </main>
</template>
