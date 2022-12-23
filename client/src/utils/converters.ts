export const camelize = (str: string) => {
  return str.replace(/(?:^\w|[A-Z]|\b\w)/g, function (word: string, index: number) {
    return index === 0 ? word.toLowerCase() : word.toUpperCase();
  }).replace(/\s+/g, '');
}

export const uncamelize = (str: string) => {
  const result = str.replace(/([A-Z])/g, " $1");
  return result.charAt(0).toUpperCase() + result.slice(1);
}