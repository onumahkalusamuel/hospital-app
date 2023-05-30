<script lang="ts" >
export default {
  name: 'Dashboard',
}; </script>;

<template>
  <div class="w-100 toolbar">
    <div>
      <a class="header-icon-link">menu</a>
      <a class="header-icon-link">
        <strong>Federal Medical Center</strong>
      </a>
    </div>
    <div style="flex-grow: 1;max-width: 400px;">
      <fluent-text-field class="w-100" placeholder="Search">
        <i slot="start">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-clipboard-fill" viewBox="0 0 16 16">
            <path fill-rule="evenodd" d="M10 1.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-1Zm-5 0A1.5 1.5 0 0 1 6.5 0h3A1.5 1.5 0 0 1 11 1.5v1A1.5 1.5 0 0 1 9.5 4h-3A1.5 1.5 0 0 1 5 2.5v-1Zm-2 0h1v1A2.5 2.5 0 0 0 6.5 5h3A2.5 2.5 0 0 0 12 2.5v-1h1a2 2 0 0 1 2 2V14a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3.5a2 2 0 0 1 2-2Z"/>
          </svg>
        </i>
      </fluent-text-field>
    </div>
    <div>
      <a class="header-icon-link">notif</a>
      <a class="header-icon-link">settings</a>
      <a class="header-icon-link">profile</a>
    </div>
  </div>
  <div class="" style="padding: 20px">
    <div class="heading mb-4 mt-4">
      <p>Welcome back, Dr. Hilda Baci</p>
    </div>
    <fluent-divider></fluent-divider>
    <div class="body">icons</div>
  </div>
</template>

<style>
.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  box-sizing: border-box;
  background-color: var(--accent-fill-rest);
}
.header-icon-link {
  padding: 5px;
  text-decoration: none;
  color:white;
  cursor:pointer;
}
.header-icon-link:hover {
  color:bisque;
}
</style>