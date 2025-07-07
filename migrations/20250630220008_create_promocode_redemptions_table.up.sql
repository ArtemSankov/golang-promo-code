CREATE TABLE promocode_redemptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    promocode_id UUID REFERENCES promocodes(id) ON DELETE CASCADE,
    user_id TEXT,
    redeemed_at TIMESTAMPTZ DEFAULT now()
);
