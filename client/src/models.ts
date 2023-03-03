export interface Ingredient extends IngredientDetails{
  name: string,
}

export interface IngredientDetails {
  amount: number,
  unit: string
}

export interface Recipe {
  title: string
  slug: string
  description: string
  ingredients: Ingredient[]
}