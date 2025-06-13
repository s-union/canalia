import type { Meta, StoryObj } from '@storybook/nextjs';

interface TypographyProps {
  variant: 'h1' | 'h2' | 'h3' | 'h4' | 'p' | 'blockquote';
  children: React.ReactNode;
}

const TypographyComponent: React.FC<TypographyProps> = ({ variant, children }) => {
  const Tag = variant;
  return <Tag>{children}</Tag>;
};

const meta: Meta = {
  component: TypographyComponent,
}

export default meta;

type Story = StoryObj<typeof meta>;

export const H1: Story = {
  args: {
    variant: 'h1',
    children: '見出し1'
  },
};

export const H2: Story = {
  args: {
    variant: 'h2',
    children: '見出し2',
  },
};

export const H3: Story = {
  args: {
    variant: 'h3',
    children: '見出し3',
  },
};

export const H4: Story = {
  args: {
    variant: 'h4',
    children: '見出し4',
  },
};

export const P: Story = {
  args: {
    variant: 'p',
    children: '山路を登りながら、こう考えた。智に働けば角が立つ。情に棹させば流される。意地を通せば窮屈だ。とかくに人の世は住みにくい。住みにくさが高じると、安い所へ引き越したくなる。どこへ越しても住みにくいと悟った時、詩が生れて、画が出来る。人の世を作ったものは神でもなければ鬼でもない。やはり向う三軒両隣りにちらちらするただの人である。ただの人が作った人の世が住みにくいからとて、越す国はあるまい。あれば人でなしの国へ行くばかりだ。人でなしの国は人の世よりもなお住みにくかろう。',
  }
}

export const Blockquote: Story = {
  args: {
    variant: 'blockquote',
    children: '山路を登りながら、こう考えた。智に働けば角が立つ。情に棹させば流される。意地を通せば窮屈だ。とかくに人の世は住みにくい。',
  }
}