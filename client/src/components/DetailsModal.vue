<script setup lang="ts">
import axios from 'axios'
import { onMounted, ref, watch } from 'vue'
import { ApiResponse, ExtendedRecipe } from '../models'
import { recipeStore } from '../stores/recipe.store'
import IconElement from './Icons/IconElement.vue'
import RecipeForm from './RecipeForm.vue'

const props = defineProps<{ recipe: ExtendedRecipe | null }>()
const emit = defineEmits(['modal-close', 'recipe-remove'])
const baseRefHidden = ref<boolean>()
const dialogRefHidden = ref<boolean>()
const dialogBgRefHidden = ref<boolean>()
const isEdit = ref<boolean>()

onMounted(() => {
  baseRefHidden.value = true
})

const isOpen = !!props.recipe

watch(
  () => props.recipe,
  () => {
    if (props.recipe) {
      baseRefHidden.value = false
      dialogRefHidden.value = false
      dialogBgRefHidden.value = false
    } else {
      baseRefHidden.value = true
      dialogRefHidden.value = true
      dialogBgRefHidden.value = true
      recipeStore.clear()
      isEdit.value = false
    }
  }
)

const handleEdit = () => {
  recipeStore.setIngredients(props.recipe?.ingredients)
  isEdit.value = true
}

const handleRemove = () => {
  if (confirm(`Are you sure you want to remove "${props.recipe?.title}"?`)) {
    axios.delete<ApiResponse>(`${import.meta.env.VITE_BACKEND_URL}/recipe/${props.recipe?.id}`).then((resp) => {
      if (resp.data.success) {
        if (props.recipe) {
          emit('recipe-remove')
          emit('modal-close')
        }
      }
    })
  }
}

const handleSuccess = (title: string, description: string) => {
  axios
    .post<ApiResponse>(`${import.meta.env.VITE_BACKEND_URL}/recipe?id=${props.recipe?.id}`, {
      title: title,
      description: description,
      ingredients: recipeStore.ingredients,
    })
    .then((resp) => {
      if (resp.data.success) {
        const url = window.location.href
        window.location.href = `${url}${url.indexOf('?') > -1 ? '&' : '?'}success=true`
      }
    })
    .catch((err) => console.error(err))
    .finally(() => {
      isEdit.value = false
      recipeStore.clear()
    })
}
</script>
<template>
  <div :class="isOpen && 'open'" class="base" :hidden="baseRefHidden" v-if="!!recipe">
    <div class="dialog-bg" @click="$emit('modal-close')" tabindex="-1" :hidden="dialogBgRefHidden"></div>
    <div
      class="dialog"
      role="dialog"
      aria-modal="true"
      :aria-hidden="isOpen ? 'false' : 'true'"
      aria-labelledby="title"
      tabindex="0"
      :hidden="dialogRefHidden"
    >
      <header>
        <h2 id="title">
          {{ recipe?.title }}
        </h2>
        <div>
          <button class="close" label="close" @click="$emit('modal-close')">x</button>
        </div>
      </header>
      <div class="body">
        <recipe-form
          v-if="isEdit"
          :recipe="recipe"
          @on-success="handleSuccess"
          :submitText="`Edit &quot;${props.recipe?.title}&quot;`"
        ></recipe-form>
        <template v-else>
          <div class="description">{{ recipe.description }}</div>
          <div class="box">
            <ul class="ingredients">
              <li v-for="ingredient in recipe.ingredients" :key="ingredient.name">
                {{ ingredient.amount }} {{ ingredient.unit }}
                {{ ingredient.name }}
              </li>
            </ul>
          </div>
          <div class="button-group">
            <button class="edit" @click="handleEdit()" :title="`Edit ${recipe.title}`">
              <icon-element name="edit" /><span>Edit</span>
            </button>
            <button class="remove" @click="handleRemove()" :title="`Remove ${recipe.title}`">
              <icon-element name="remove" /><span>Delete</span>
            </button>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>
<style scoped>
[hidden] {
  display: none !important;
}

template {
  display: contents;
  width: 30rem;
}

header {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #36363b;
  padding: 1.5rem 1rem;
}

.body {
  padding: 1.5rem 1rem;
  overflow-y: auto;
}

.dialog {
  z-index: 1;
  display: flex;
  flex-direction: column;
  background-color: #242428;
  box-shadow: 1px 1px 1px #0000002e;
  border-radius: 2px;
  width: 30rem;
  max-height: 90vh;
}

.base.open .dialog {
  display: flex;
  opacity: 1;
}

.base {
  position: fixed;
  inset: 0;
  z-index: 2;
  margin: 2rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.dialog-bg {
  position: fixed;
  inset: 0px;
  background-color: rgba(0, 0, 0, 0.8);
}

.close {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem 1rem;
}

h2 {
  display: block;
  margin: 0;
}

.description {
  white-space: pre-line;
}

.button-group {
  display: flex;
  width: 100%;
  gap: 0.5rem;
  align-items: center;
}

.button-group button {
  height: 2.5rem;
  min-width: 4rem;
  flex: 1;
  gap: 0.5rem;
  font-size: 0.75rem;
  white-space: nowrap;
}

.edit {
  background-color: var(--color-blue-fade);
}

.remove {
  background-color: var(--color-orange-fade);
}

.submit {
  background-color: var(--color-green-fade);
}
</style>
