<script setup lang="ts">
import { ref } from "vue";
import AddIngredients from "./AddIngredients.vue";
import { recipeStore } from "../../stores/recipe.store";
import axios from "axios";

const title = ref<string>("");
const description = ref<string>("");
const errors = ref<string[]>([]);

const handleSubmit = () => {
  errors.value = [];

  if (!title.value) {
    errors.value.push("Title cannot be empty!");
  }

  if (!description.value) {
    errors.value.push("Description cannot be empty!");
  }

  const hasAnyIngredientField = recipeStore.ingredients.some(
    (ingredient) => !(ingredient.amount && ingredient.name && ingredient.unit)
  );

  if (hasAnyIngredientField) {
    errors.value.push("Ingredients field cannot be empty!");
  }

  if (!description.value || hasAnyIngredientField || !title.value) {
    return;
  }

  axios
    .post(`${import.meta.env.VITE_BACKEND_URL}/recipe`, {
      title: title.value,
      description: description.value,
      ingredients: recipeStore.ingredients,
    })
    .catch((err) => console.error(err));
};
</script>

<template>
  <div>
    <h1>Add recipe</h1>
    <form class="add-form" v-on:submit.prevent="handleSubmit">
      <div class="form-section" style="background: #647dee44">
        <label class="form-group"
          >Title<input type="text" name="ingredient-title" v-model="title"
        /></label>
      </div>
      <div class="form-section" style="background: #647dee44">
        <label class="form-group"
          >Description<textarea
            v-model="description"
            rows="4"
            name="recipe-step"
          ></textarea>
        </label>
      </div>
      <add-ingredients />

      <div class="form-section" style="background-color: #aff6cf44">
        <div v-if="errors.length">
          <b>Please correct the following error(s):</b>
          <ul class="errors">
            <li v-for="(error, index) in errors" :key="index">{{ error }}</li>
          </ul>
        </div>
        <button class="submit">Add recipe!</button>
      </div>
    </form>
  </div>
</template>

<style scoped>
.form-section {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding: 1rem 0.5rem;
  border-radius: var(--border-radius);
}

.form-group {
  display: flex;
  gap: 1rem;
  flex-direction: column;
  align-items: baseline;
}

form {
  margin: auto;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

:deep() input,
:deep() select,
:deep() textarea {
  box-sizing: border-box;
  width: 100%;
  padding: 0.5rem 1rem;
}

:deep() button.full-width {
  width: 100%;
}

button.submit {
  background: none;
  border: 3px solid;
}

.errors {
  text-align: left;
}
</style>
