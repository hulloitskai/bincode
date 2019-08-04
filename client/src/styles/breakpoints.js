export const screenSizes = {
  DESKTOP: 1500,
  LAPTOP: 1200,
  TABLET: 900,
  PHABLET: 700,
  MOBILE: 480,
};

export const getScreenSize = () =>
  window.innerWidth ||
  document.documentElement.clientWidth ||
  document.body.clientWidth;
