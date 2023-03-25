<script setup lang="ts">
import { ref } from 'vue'
import { Recipe } from '../models'
import { recipeStore } from '../stores/recipe.store'
import AddIngredient from './AddRecipe/AddIngredient.vue'

const props = defineProps<{ recipe?: Recipe; submitText?: string }>()
const emit = defineEmits(['on-success'])
const title = ref<string>(props.recipe?.title || '')
const description = ref<string>(props.recipe?.description || '')
const errors = ref<string[]>([])
const handleSubmit = () => {
  errors.value = []

  if (!title.value) {
    errors.value.push('Title cannot be empty!')
  }

  if (!description.value) {
    errors.value.push('Description cannot be empty!')
  }

  const hasAnyIngredientField = recipeStore.ingredients.some(
    (ingredient) => !(ingredient.amount && ingredient.name && ingredient.unit)
  )

  if (hasAnyIngredientField) {
    errors.value.push('Ingredients field cannot be empty!')
  }

  if (!description.value || hasAnyIngredientField || !title.value) {
    return
  }

  emit('on-success', title.value, description.value)
}
</script>
<template>
  <form class="add-form" v-on:submit.prevent="handleSubmit">
    <div class="form-section" style="background: #647dee44">
      <label class="form-group">
        <span class="title">Title</span>
        <input type="text" name="ingredient-title" v-model="title" />
      </label>
    </div>
    <div class="form-section" style="background: #647dee44">
      <label class="form-group">
        <span class="title">Description</span>
        <textarea v-model="description" rows="4" name="recipe-step"></textarea>
      </label>
    </div>
    <div class="form-section" style="background: #9fa4c444">
      <span class="title">Ingredients</span>
      <div v-for="(ingredient, index) in recipeStore.ingredients" :key="index" class="form-group ingredient">
        <add-ingredient :ingredient="ingredient" :index="index" />
      </div>
      <button
        v-if="recipeStore.ingredients[recipeStore.ingredients.length - 1].name !== ''"
        class="full-width"
        type="button"
        @click="recipeStore.addIngredient()"
      >
        Add ingredient
      </button>
    </div>
    <div class="form-section" style="background-color: #aff6cf44">
      <div v-if="errors.length">
        <b>Please correct the following error(s):</b>
        <ul class="errors">
          <li v-for="(error, index) in errors" :key="index">{{ error }}</li>
        </ul>
      </div>
      <button>{{ submitText || 'Add recipe!' }}</button>
    </div>
  </form>
</template>
<style scoped>
.form-section {
  display: flex;
  flex-direction: column;
  padding: 1rem 0.5rem;
  border-radius: var(--border-radius);
}

.form-group {
  display: flex;
  flex-direction: column;
  align-items: baseline;
}

.ingredient {
  flex-direction: row;
  display: flex;
  gap: 0.5rem;
  align-items: end;
  padding: 0.5rem;
  background: #1c8c7344;
  border-radius: 4px;
  margin-bottom: 0.5rem;
}

.title {
  font-size: 1.25rem;
  margin-bottom: 0.5rem;
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

.errors {
  text-align: left;
}
</style>
