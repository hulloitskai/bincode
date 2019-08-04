/**
 * Sets a timeout relative to another time.
 *
 * @param {Date} time
 * @param {TimerHandler} handler
 * @param {number} timeout
 * @param {any[]} arguments
 */
export const setTimeoutFrom = (time, handler, timeout, ...args) => {
  const delay = timeout - (new Date().getTime() - time.getTime());
  if (delay > 0) setTimeout(handler, delay, ...args);
  else handler(...args);
};
