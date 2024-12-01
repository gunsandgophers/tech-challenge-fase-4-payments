Feature: Create new Payment
  I need to request a link to new payment

  Scenario: Create a new payment
    Given That I need to create a new payment via API
    And have the order ID "c8f32477-404b-4db1-bf2e-23a3c0d8f6b0"
    And amount as "30.5" 
    When I send the data
    Then the payment need to be created with link and order ID "c8f32477-404b-4db1-bf2e-23a3c0d8f6b0"
