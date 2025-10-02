-- Remove seed data for cost_tracking table
DELETE FROM cost_trackings WHERE planting_cycle_id = '550e8400-e29b-41d4-a716-446655440001'::uuid;
