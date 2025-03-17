-- Crear extensiones útiles de PostgreSQL
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "pg_trgm"; -- Para búsquedas de texto eficientes

-- Tabla de Usuarios
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    date_of_birth DATE,
    profile_picture_url VARCHAR(255),
    is_verified BOOLEAN DEFAULT FALSE,
    verification_token UUID,
    reset_password_token UUID,
    reset_password_expires TIMESTAMP,
    last_login TIMESTAMP,
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger para mantener updated_at actualizado
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_modtime
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Organizadores
CREATE TABLE organizers (
    organizer_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    company_name VARCHAR(255) NOT NULL,
    company_description TEXT,
    logo_url VARCHAR(255),
    website VARCHAR(255),
    tax_id VARCHAR(50),
    address TEXT,
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    is_verified BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) DEFAULT 'pending', -- pending, active, suspended
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_organizers_modtime
    BEFORE UPDATE ON organizers
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Categorías de Eventos
CREATE TABLE event_categories (
    category_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    icon_url VARCHAR(255),
    slug VARCHAR(100) UNIQUE,
    color_hex VARCHAR(7),
    is_active BOOLEAN DEFAULT TRUE,
    display_order INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_categories_modtime
    BEFORE UPDATE ON event_categories
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Lugares (Venues)
CREATE TABLE venues (
    venue_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    country VARCHAR(100) NOT NULL,
    postal_code VARCHAR(20),
    -- Usando PostGIS para ubicación
    location GEOGRAPHY(POINT),
    capacity INTEGER,
    description TEXT,
    directions TEXT,
    image_url VARCHAR(255),
    website VARCHAR(255),
    phone VARCHAR(20),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_venues_modtime
    BEFORE UPDATE ON venues
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Eventos
CREATE TABLE events (
    event_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organizer_id UUID NOT NULL REFERENCES organizers(organizer_id) ON DELETE CASCADE,
    category_id UUID REFERENCES event_categories(category_id),
    venue_id UUID REFERENCES venues(venue_id),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE,
    description TEXT,
    short_description VARCHAR(500),
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    doors_open_date TIMESTAMP WITH TIME ZONE,
    is_published BOOLEAN DEFAULT FALSE,
    is_featured BOOLEAN DEFAULT FALSE,
    is_private BOOLEAN DEFAULT FALSE,
    main_image_url VARCHAR(255),
    status VARCHAR(20) DEFAULT 'upcoming', -- upcoming, ongoing, completed, cancelled
    ticket_sale_start TIMESTAMP WITH TIME ZONE,
    ticket_sale_end TIMESTAMP WITH TIME ZONE,
    event_url VARCHAR(255),
    terms_and_conditions TEXT,
    max_tickets_per_person INTEGER,
    tags TEXT[],
    metadata JSONB,
    seo_title VARCHAR(255),
    seo_description VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_events_modtime
    BEFORE UPDATE ON events
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Índice para búsqueda de texto en eventos
CREATE INDEX events_title_description_idx ON events USING GIN ((title || ' ' || COALESCE(description, '')) gin_trgm_ops);
CREATE INDEX events_tags_idx ON events USING GIN (tags);

-- Tabla de Imágenes de Eventos
CREATE TABLE event_images (
    image_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id UUID NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
    image_url VARCHAR(255) NOT NULL,
    alt_text VARCHAR(255),
    is_main BOOLEAN DEFAULT FALSE,
    display_order INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Tipos de Entradas
CREATE TABLE ticket_types (
    ticket_type_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id UUID NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    original_price DECIMAL(10, 2), -- Para mostrar descuentos
    quantity INTEGER NOT NULL,
    sold_count INTEGER DEFAULT 0,
    reserved_count INTEGER DEFAULT 0,
    max_per_purchase INTEGER,
    min_per_purchase INTEGER DEFAULT 1,
    start_sale_date TIMESTAMP WITH TIME ZONE,
    end_sale_date TIMESTAMP WITH TIME ZONE,
    is_active BOOLEAN DEFAULT TRUE,
    requires_approval BOOLEAN DEFAULT FALSE,
    is_hidden BOOLEAN DEFAULT FALSE,
    sort_order INTEGER,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_ticket_types_modtime
    BEFORE UPDATE ON ticket_types
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Promociones
CREATE TABLE promotions (
    promotion_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organizer_id UUID REFERENCES organizers(organizer_id) ON DELETE CASCADE,
    code VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255),
    discount_type VARCHAR(20) NOT NULL, -- percentage, fixed
    discount_value DECIMAL(10, 2) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE,
    max_uses INTEGER,
    current_uses INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    min_purchase_amount DECIMAL(10, 2) DEFAULT 0,
    max_discount_amount DECIMAL(10, 2),
    applies_to_all_events BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_promotions_modtime
    BEFORE UPDATE ON promotions
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Promociones por Evento (solo necesaria si applies_to_all_events es FALSE)
CREATE TABLE event_promotions (
    event_promotion_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id UUID NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
    promotion_id UUID NOT NULL REFERENCES promotions(promotion_id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(event_id, promotion_id)
);

-- Tabla de Órdenes
CREATE TABLE orders (
    order_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(user_id) ON DELETE SET NULL,
    event_id UUID NOT NULL REFERENCES events(event_id),
    order_reference VARCHAR(100) NOT NULL UNIQUE,
    total_amount DECIMAL(10, 2) NOT NULL,
    subtotal_amount DECIMAL(10, 2) NOT NULL,
    discount_amount DECIMAL(10, 2) DEFAULT 0.00,
    service_fee DECIMAL(10, 2) DEFAULT 0.00,
    net_amount DECIMAL(10, 2) NOT NULL,
    promotion_id UUID REFERENCES promotions(promotion_id),
    status VARCHAR(20) NOT NULL, -- pending, completed, cancelled, refunded
    payment_status VARCHAR(20) DEFAULT 'pending', -- pending, authorized, paid, failed, refunded
    payment_method VARCHAR(50),
    payment_reference VARCHAR(255),
    payment_gateway VARCHAR(50),
    checkout_expires_at TIMESTAMP WITH TIME ZONE,
    ip_address VARCHAR(45),
    user_agent TEXT,
    notes TEXT,
    tracking_data JSONB,
    currency VARCHAR(3) DEFAULT 'USD',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_orders_modtime
    BEFORE UPDATE ON orders
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Detalles de Órdenes (Entradas compradas)
CREATE TABLE order_items (
    order_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    ticket_type_id UUID NOT NULL REFERENCES ticket_types(ticket_type_id),
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Entradas (Boletos individuales)
CREATE TABLE tickets (
    ticket_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_item_id UUID NOT NULL REFERENCES order_items(order_item_id) ON DELETE CASCADE,
    event_id UUID NOT NULL REFERENCES events(event_id),
    ticket_reference VARCHAR(100) NOT NULL UNIQUE,
    qr_code VARCHAR(255),
    barcode VARCHAR(255),
    check_in_status BOOLEAN DEFAULT FALSE,
    check_in_time TIMESTAMP WITH TIME ZONE,
    check_in_notes TEXT,
    check_in_by UUID REFERENCES users(user_id),
    attendee_name VARCHAR(255),
    attendee_email VARCHAR(255),
    attendee_phone VARCHAR(50),
    is_void BOOLEAN DEFAULT FALSE,
    is_transferable BOOLEAN DEFAULT TRUE,
    transfer_token UUID,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_tickets_modtime
    BEFORE UPDATE ON tickets
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Datos de Facturación
CREATE TABLE billing_information (
    billing_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    full_name VARCHAR(255) NOT NULL,
    tax_id VARCHAR(50),
    company_name VARCHAR(255),
    address_line1 VARCHAR(255),
    address_line2 VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_billing_modtime
    BEFORE UPDATE ON billing_information
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Transferencias a Organizadores
CREATE TABLE organizer_payouts (
    payout_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organizer_id UUID NOT NULL REFERENCES organizers(organizer_id) ON DELETE CASCADE,
    event_id UUID REFERENCES events(event_id),
    amount DECIMAL(10, 2) NOT NULL,
    transaction_fee DECIMAL(10, 2) DEFAULT 0.00,
    status VARCHAR(20) NOT NULL, -- pending, processing, completed, failed, cancelled
    payout_date TIMESTAMP WITH TIME ZONE,
    bank_account_id UUID,
    reference VARCHAR(255),
    notes TEXT,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_payouts_modtime
    BEFORE UPDATE ON organizer_payouts
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Configuración de Comisiones
CREATE TABLE commission_settings (
    commission_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organizer_id UUID REFERENCES organizers(organizer_id) ON DELETE CASCADE,
    event_id UUID REFERENCES events(event_id) ON DELETE CASCADE,
    commission_percentage DECIMAL(5, 2) NOT NULL,
    service_fee_percentage DECIMAL(5, 2) NOT NULL,
    payment_gateway_fee_percentage DECIMAL(5, 2) DEFAULT 0.00,
    fixed_fee DECIMAL(5, 2) DEFAULT 0.00,
    is_active BOOLEAN DEFAULT TRUE,
    applies_to_all_events BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_commission_relation CHECK (
        (organizer_id IS NOT NULL AND applies_to_all_events = TRUE) OR
        (event_id IS NOT NULL AND applies_to_all_events = FALSE)
    )
);

CREATE TRIGGER update_commission_modtime
    BEFORE UPDATE ON commission_settings
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Tabla de Bancos para Transferencias
CREATE TABLE bank_accounts (
    bank_account_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organizer_id UUID NOT NULL REFERENCES organizers(organizer_id) ON DELETE CASCADE,
    bank_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(50) NOT NULL,
    account_holder VARCHAR(255) NOT NULL,
    account_type VARCHAR(50),
    routing_number VARCHAR(50),
    branch_name VARCHAR(255),
    branch_code VARCHAR(50),
    swift_code VARCHAR(50),
    iban VARCHAR(50),
    is_primary BOOLEAN