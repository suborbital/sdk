import { log } from "@suborbital/plugin";

export const run = (input) => {
  let message = "Hello, " + input;

  log.info(message);

  return message;
};
