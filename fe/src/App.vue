<script setup>
import EditorPopup from '@/components/EditorPopup.vue'

const lipsum = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed dictum ex ac mollis suscipit. Suspendisse nec molestie massa. Sed urna enim, dictum eu sodales in, sollicitudin non urna. Curabitur vehicula ligula non metus fermentum tempor. Maecenas turpis orci, cursus interdum bibendum et, consequat sed est. Quisque scelerisque, ante vel posuere maximus, lacus ante fermentum lectus, a ultricies est metus nec diam. In in egestas lorem, quis cursus mi. Proin purus dolor, aliquet consectetur eros non, dignissim posuere leo. Sed id neque diam. In facilisis nunc nisi, eget ultrices massa laoreet a.\n\nSed tempus arcu eget eros commodo lacinia. Aenean vel orci a nisi rutrum ornare. Phasellus pharetra sem eget nibh interdum, in volutpat eros ultrices. Fusce tempor, erat vitae pretium blandit, turpis mi malesuada tortor, in dictum arcu velit non ante. Duis et purus facilisis, porta tortor quis, iaculis diam. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin cursus placerat suscipit. Duis neque augue, faucibus a erat a, interdum egestas turpis. Ut at auctor lorem.\n\nVestibulum accumsan elit ut lectus sagittis dignissim. Suspendisse sit amet dapibus dui, aliquet volutpat urna. Phasellus quis dolor quis dui pretium imperdiet. Phasellus commodo leo quam, non consectetur justo mattis sit amet. Fusce tempus imperdiet odio, vitae cursus sem. Suspendisse ultrices risus eu mauris ultricies, eget scelerisque nisl vulputate. Vivamus ultricies dolor ut elementum accumsan. Vivamus molestie justo nunc, et euismod mi lacinia non.\n\nPhasellus ultricies interdum urna, vitae sodales metus semper vel. Etiam odio purus, placerat in nibh et, cursus interdum odio. Maecenas egestas dolor quis ipsum viverra malesuada. Suspendisse venenatis molestie maximus. Duis sagittis, neque sed interdum sagittis, turpis magna mattis purus, et sodales eros dui vel dui. Nunc vitae ligula in tortor dapibus laoreet eget ut turpis. Ut pellentesque tristique feugiat. Quisque fringilla quam ac ligula mattis, varius tincidunt tortor pharetra. Nullam bibendum nisl eget sapien faucibus, at consequat ipsum ornare. Phasellus leo nulla, mollis eu tristique nec, bibendum et lacus. Phasellus pellentesque nibh ac aliquam euismod. Ut vestibulum velit a laoreet porttitor. Vivamus odio tortor, vehicula ut scelerisque ac, pellentesque id elit. Fusce enim quam, rhoncus ac eros vel, blandit dapibus tortor.\n\nVivamus ultricies sagittis erat ut rutrum. Aenean egestas dignissim lectus. Nunc varius, nunc quis pellentesque faucibus, odio lectus porttitor ipsum, nec dictum magna justo et sapien. Nulla facilisi. Duis vel lobortis urna. Suspendisse rhoncus mi ac eros ultrices pharetra. Aliquam eget erat eget purus blandit sagittis. Vivamus placerat erat justo, sit amet fermentum felis elementum in. Sed sed sapien eget justo vulputate sagittis in ut sem. Duis viverra tortor eu convallis accumsan. Ut ut lectus non ante vestibulum ullamcorper. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Aenean volutpat metus sit amet libero consequat interdum. Phasellus quis accumsan massa.'
</script>

<template>
  <main class="font-raleway p-8">
    <h1 class="text-center mb-8 text-2xl cursor-pointer">My Dictionary</h1>
    <div class="flex flex-col w-[1024px] m-auto border-2 rounded-xl">

      <div class="headers w-full flex flex-row">
        <div class="w-[270px] px-4 py-2">
          <input type="text" class="border-[1px] ml-3 rounded-[24px] text-lg px-4 py-2"
            placeholder="Type to search.." />
        </div>
        <div class="grow px-4 py-2 flex items-center">Description</div>
      </div>

      <div class="body w-full flex flex-row h-[500px]">
        <div class="w-[270px] grow-0 shrink-0 basis-[270px] px-4 py-2 overflow-y-auto h-full">
          <ul>
            <li v-for="i in 100" class="px-2 py-1 hover:bg-[#f2f2f2] cursor-pointer">Term {{ i }}</li>
          </ul>
        </div>
        <div class="shrink p-4 overflow-y-auto h-full">
          <div>
            {{ lipsum }}
          </div>
        </div>
      </div>

    </div>

    <EditorPopup></EditorPopup>
  </main>
</template>
