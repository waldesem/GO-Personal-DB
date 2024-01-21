const server = 'http://localhost:5000';
//const server = '';

function debounce(func: (...args: any[]) => void, delay: number): (...args: any[]) => void {

  let timer: ReturnType<typeof setTimeout> | undefined;

  return function(this: any, ...args: any[]) {
    if (timer) clearTimeout(timer);

    timer = setTimeout(() => {
      func.apply(this, args);
      timer = undefined; // Reset the timer after the function is called
    }, delay);
  };
};

function clearForm(item: Object): void {
  Object.keys(item).forEach(key => {
    delete item[key as keyof typeof item];
  });
};

function timeSince(date: string) {

  const seconds: number = Math.floor(((new Date() as any) - (new Date(date) as any)) / 1000);

  let interval = seconds / 31536000;

  if (interval > 1) {
    return Math.floor(interval) + " лет назад";
  };
  interval = seconds / 2592000;
  if (interval > 1) {
    return Math.floor(interval) + " месяцев назад";
  };
  interval = seconds / 86400;
  if (interval > 1) {
    return Math.floor(interval) + " дней назад";
  };
  interval = seconds / 3600;
  if (interval > 1) {
    return Math.floor(interval) + " часов назад";
  };
  interval = seconds / 60;
  if (interval > 1) {
    return Math.floor(interval) + " минут назад";
  };
  return Math.floor(seconds) + " секунд назад";
};

export { 
  server, 
  debounce,
  timeSince,
  clearForm 
};
