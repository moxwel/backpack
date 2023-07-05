// design.sv

// Memoria
module ROM(input  logic [5:0] addr,
           output logic [18:0] data);
  always @(*)
  begin
    // mem => arreglo de 64 espacios
    // cada espacio almacena 19 bits
    logic [18:0] mem [0:63];

    // Definicion de ROM
    mem[0] = 19'h01713;
    mem[1] = 19'h1034C;
    mem[2] = 19'h1074C;
    mem[3] = 19'h21F05;
    mem[4] = 19'h31F02;
    mem[5] = 19'h45D52;
    mem[6] = 19'h55D52;
    mem[7] = 19'h65D52;
    mem[8] = 19'h75D52;
    
    ////////////////////////////////////
    // Definir mas instrucciones AQUI //
    ////////////////////////////////////
    
    // mem[9] = 19'hFFFFF;
    // ...

    assign data = mem[addr];
  end
endmodule


// Splitter
module SPLITTER(input  logic [18:0] in,
                output logic [3:0]  op,
                output logic [7:0]  a,
                output logic [7:0]  b);
  always @(*)
  begin
    assign op = in[18:16]; // Primeros 3 bits
    assign a = in[15:8];   // Luego 8 bits
    assign b = in[7:0];    // Ultimos 8 bits
  end
endmodule


// ALU
module ALU(input  logic [2:0] op,
           input  logic [7:0] a,
           input  logic [7:0] b,
           output logic [7:0] result);
  always @(*)
  begin
    case(op)
      3'b000: // Suma
        begin
          logic [8:0] carry;
          carry = 0;
          for (int i = 0; i < 8; i++) begin
            result[i] = a[i] ^ b[i] ^ carry[i];
            carry[i+1] = (a[i] & b[i]) | (carry[i] & (a[i] ^ b[i]));
          end
        end


      3'b001: // Resta
        begin
          logic [8:0] borrow;
          borrow = 0;
          for (int i = 0; i < 8; i++) begin
            result[i] = a[i] ^ b[i] ^ borrow[i];
            borrow[i+1] = (~a[i] & b[i]) | (borrow[i] & ~(a[i] ^ b[i]));
          end
        end

      3'b010: assign result = a << b; // Bitshift izq
      3'b011: assign result = a >> b; // Bitshift der

      3'b100: assign result = a & b;  // Bitwise Composicion
      3'b101: assign result = a | b;  // Bitwise Disyuncion
      3'b110: assign result = a ^ b;  // Bitwise Exclusion

      3'b111: assign result = ~a;     // Negacion
    endcase
  end
endmodule
