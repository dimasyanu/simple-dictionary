<script setup>
import { ref } from 'vue';
const state = ref({
  id: '',
  key: '',
  description: ''
})
</script>

<template>
  <div>
    <div class="overlay fixed top-0 left-0 w-full h-full bg-[#00000080] z-[1]">
      <div class="flex flex-col items-center justify-center w-full h-full">
        <div class="min-h-[24rem] w-[50rem] bg-white rounded-xl p-8">
          <div>
            <input type="text"  class="border-[1px] rounded-[24px] text-lg px-4 py-2 w-full text-center" placeholder="Type to input key.." v-model="state.key"/>
          </div>
          <div class="mt-4">
            <textarea name="description" id="description" class="border-[1px] w-full min-h-[240px] rounded-xl p-4" placeholder="Type to input description.." v-model="state.description"></textarea>
          </div>
          <div class="text-center">
            <div class="inline-block px-5 py-2 cursor-pointer bg-[#42b883] rounded-md text-white mt-3">Add</div>
            <div class="inline-block px-5 py-2 cursor-pointer bg-[#8b8b8b] rounded-md text-white ms-3 mt-3">Cancel</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>