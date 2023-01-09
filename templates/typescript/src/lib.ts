import { log } from "@suborbital/plugin";

export const run = (input: string): string => {
  let message = "Hello, " + input;

  log.info(message);

  return message;
};
