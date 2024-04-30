<script setup>
import { computed, ref, watch } from 'vue';
import EditorPopup from '@/components/EditorPopup.vue'

const state = ref({
  current: {
    id: null,
    key: null,
    description: null
  },
  keyword: '',
  items: [
    {
      id: '4b29ba3f-4904-445e-8b6b-96539f3f457d',
      key: 'Lorem ipsum',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam cursus mi non semper cursus. Nullam sed elit vitae dui condimentum feugiat. Vivamus tincidunt sit amet augue vitae dapibus. Quisque eget ante urna. Cras cursus a nibh in interdum. In hac habitasse platea dictumst. Suspendisse sit amet nisl dignissim, viverra lectus ac, hendrerit nisi. Fusce quis egestas libero. Aliquam urna quam, vulputate non maximus sit amet, eleifend id eros. Suspendisse vitae sagittis quam. Mauris non nulla enim. Vivamus leo risus, tincidunt non ligula nec, ultricies cursus odio. In hac habitasse platea dictumst. Fusce vel tortor lorem. Integer id mattis ipsum. Suspendisse vel lobortis turpis.'
    },
    {
      id: 'a148e645-dcfa-4153-9401-4f6bb03ccb08',
      key: 'Integer pulvinar',
      description: 'Integer pulvinar laoreet mi sit amet pretium. Suspendisse blandit fringilla tortor nec rutrum. Curabitur eu leo tellus. Quisque tincidunt cursus diam ac aliquet. Quisque tempus, nisi sit amet vestibulum efficitur, nisi sem ullamcorper eros, nec aliquam diam est vitae massa. Suspendisse sollicitudin, ipsum quis imperdiet bibendum, dolor libero aliquet lorem, ut auctor lectus ante in tortor. Morbi non enim augue. Donec ultricies nisi nec urna consectetur convallis. Duis nibh enim, volutpat id convallis ut, semper nec metus. Aenean egestas turpis ac neque molestie pulvinar.'
    },
    {
      id: 'e54c4aa0-a405-4064-8c3a-bd86c67d389e',
      key: 'Nullam finibus',
      description: 'Nullam finibus justo in est posuere dictum eu in ligula. Nunc volutpat consectetur erat, id pretium augue eleifend id. Vestibulum non arcu dapibus, commodo eros sed, feugiat urna. Sed a commodo sapien, congue cursus risus. Proin at odio vel nisl auctor facilisis in nec erat. Suspendisse gravida mattis diam at ultricies. Maecenas eget nibh dui. Maecenas laoreet mauris non elit aliquam euismod. Sed et eros eget lorem efficitur aliquet. Nulla facilisi.'
    },
    {
      id: 'fa4c63c0-cfbd-439e-b602-1bfc760dea18',
      key: 'Aliquam cursus',
      description: 'Aliquam cursus ultrices justo, non porttitor tellus aliquam vel. Aenean et efficitur nisi. Nullam ullamcorper purus nec quam gravida, sed consectetur purus pulvinar. Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis efficitur massa quis erat luctus, quis cursus erat porta. Duis egestas ex in arcu rhoncus, at gravida nulla imperdiet. Donec mattis urna in ex varius dapibus. Praesent venenatis laoreet urna sed rutrum. Aenean laoreet erat in gravida rutrum. Etiam eget hendrerit risus, et efficitur ex. Cras hendrerit risus vel orci feugiat auctor. Donec vitae bibendum arcu. Morbi pharetra nisi orci.'
    },
    {
      id: 'fdc3c887-d37b-48b7-88b5-33bb3392c1ff',
      key: 'Maecenas fringilla',
      description: 'Maecenas fringilla felis id nunc pharetra tristique. Fusce facilisis, dui sit amet fermentum laoreet, odio lacus cursus dui, vel tristique tellus magna at velit. Ut eu eros fringilla, mollis neque eget, convallis massa. Etiam interdum non nisi sed pulvinar. Fusce placerat ex quis ultrices convallis. Aliquam aliquet scelerisque dolor, vel maximus nisi fermentum in. Ut ut lobortis quam.'
    }
  ],
  showPopup: false
})

const filteredItems = computed(() => {
  if (state.value.keyword == null || state.value.keyword === '') return state.value.items
  return state.value.items.filter(x => x.key.toLowerCase().includes(state.value.keyword.toLowerCase()))
})

const selectItem = function(id) {
  state.value.current.id = id
  let item = state.value.items.find(x => x.id === id)
  state.value.current.key = item.key
  state.value.current.description = item.description
}

const add = function() {
  state.value.current.id = null
  state.value.current.key = ''
  state.value.current.description = ''
  state.value.showPopup = true
}

const edit = function() {
  state.value.showPopup = true
}

const save = function(payload) {
  let id = payload.id
  let key = payload.key
  let desc = payload.description
  state.value.showPopup = false
}

watch(filteredItems, newVal => {
  state.value.current.id = newVal[0].id
  state.value.current.key = newVal[0].key
  state.value.current.description = newVal[0].description
})
</script>

<template>
  <main class="font-raleway p-8">
    <h1 class="text-center mb-8 text-2xl cursor-pointer" @click="add">My Dictionary</h1>
    <div class="flex flex-col w-[1024px] m-auto border-2 rounded-xl">

      <div class="headers w-full flex flex-row">
        <div class="w-[270px] px-4 py-2">
          <input type="text" class="border-[1px] ml-3 rounded-[24px] text-lg px-4 py-2"
            placeholder="Type to search.." v-model="state.keyword" />
        </div>
        <div class="grow px-4 py-2 flex items-center" >Description<span v-if="state.current.id != null" class="cursor-pointer" @click="edit"> - Edit</span></div>
      </div>

      <div class="body w-full flex flex-row h-[500px]">
        <div class="w-[270px] grow-0 shrink-0 basis-[270px] px-4 py-2 overflow-y-auto h-full">
          <ul>
            <li v-for="item in filteredItems" class="item px-2 py-1 hover:bg-[#f2f2f2] cursor-pointer" :class="{ active: item.id === state.current.id }" @click="selectItem(item.id)">{{ item.key }}</li>
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
      :term="state.current.key"
      :description="state.current.description"
      @onSave="save"
      @onCancel="state.showPopup = false" />
  </main>
</template>
