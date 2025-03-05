export function ButtonStyle (mode: 'primary' | 'secondary') {
  const baseStyle = 'px-4 py-2 rounded font-bold text-lg hover:opacity-80 hover:scale-[0.98] transition-all duration-300';

  return `${baseStyle} ${mode === 'primary' ? 'bg-primary text-foreground' : 'bg-secondary text-white'}`;
}
