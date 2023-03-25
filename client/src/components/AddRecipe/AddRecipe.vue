<script setup lang="ts">
import axios from "axios";
import { ApiResponse } from "../../models";
import { recipeStore } from "../../stores/recipe.store";
import RecipeForm from "../RecipeForm.vue";

const handleSuccess = (title: string, description: string) => {
  axios
    .post<ApiResponse>(`${import.meta.env.VITE_BACKEND_URL}/recipe`, {
      title: title,
      description: description,
      ingredients: recipeStore.ingredients,
    })
    .then((resp) => {
      if (resp.data.success) {
        const url = window.location.href;
        window.location.href = `${url}${url.indexOf("?") > -1 ? "&" : "?"
          }success=true`;
      }
    })
    .catch((err) => console.error(err)).finally(() => {
      recipeStore.clear()
    });
}
</script>

<template>
  <div>
    <h1>Add recipe</h1>
    <recipe-form @on-success="handleSuccess"></recipe-form>
  </div>
</template>

<style scoped></style>
