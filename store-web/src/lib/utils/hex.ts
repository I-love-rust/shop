import { tags } from "./colordict";

export function textToHex(text: string) {
  if (tags[text]) { return tags[text] }
  var hash = 0;
  for (var i = 0; i < text.length; i++) {
    hash = text.charCodeAt(i) + ((hash << 5) - hash);
  }
  var c = (hash & 0x00FFFFFF)
    .toString(16)
    .toUpperCase();

  return "#"+"00000".substring(0, 6 - c.length) + c;
}