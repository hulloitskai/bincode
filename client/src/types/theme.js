import includes from "lodash/includes";

const theme = {
  LIGHT: "light",
  DARK: "dark",
};

export default theme;

export const themeProp = {
  type: String,
  validator: val => includes([theme.LIGHT, theme.DARK], val),
};
