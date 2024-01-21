<script setup lang="ts">

import { server } from '@utilities/utils';

</script>

<template>
  <div class="container py-5 d-print-none">
    <footer class="d-flex justify-content-around py-3 my-4 border-top">
      <a :href="`${server}/docs`" target=”_blank” title="OpenAPI">
        <i class="bi bi-book-half"> OpenAPI</i>
      </a>
      &nbsp;
      <a href="https://github.com/waldesem/Web-Personal-DB" target=”_blank” title="Github">
        <i class="bi bi-github">  Github</i>
      </a>
    </footer>
  </div>
</template>